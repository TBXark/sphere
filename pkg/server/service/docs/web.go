package docs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
	"github.com/tbxark/sphere/pkg/server/route/cors"
	"github.com/tbxark/sphere/pkg/server/route/docs"
	"html/template"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Target struct {
	Address string
	Spec    *swag.Spec
}

type Config struct {
	Address string
	Targets []Target
}

type Web struct {
	config *Config
	server *http.Server
}

func NewWebServer(conf *Config) *Web {
	return &Web{
		config: conf,
	}
}

func (w *Web) Identifier() string {
	return "docs"
}

func (w *Web) Run() error {
	engine := gin.Default()
	cors.Setup(engine, []string{"*"})

	for _, spec := range w.config.Targets {
		if err := setup(spec.Spec, engine, spec.Address); err != nil {
			return err
		}
	}
	indexHTML := []byte(createIndex(w.config.Targets))
	engine.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", indexHTML)
	})

	w.server = &http.Server{
		Addr:    w.config.Address,
		Handler: engine.Handler(),
	}
	return w.server.ListenAndServe()
}

func (w *Web) Clean() error {
	if w.server != nil {
		return w.server.Close()
	}
	return nil
}

func setup(spec *swag.Spec, router gin.IRouter, target string) error {
	targetURL, err := url.Parse(target)
	if err != nil {
		return fmt.Errorf("invalid target URL: %v", err)
	}

	route := router.Group("/" + strings.ToLower(spec.InstanceName()))

	spec.Host = ""
	spec.BasePath = fmt.Sprintf("/%s/api", route.BasePath())
	spec.Description = spec.Description + fmt.Sprintf(" | proxy for %s", target)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	docs.Setup(route.Group("/doc"), spec)
	route.Any("/api/*path", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("path")
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	return nil
}

func createIndex(targets []Target) string {
	const indexHTML = `<!DOCTYPE html>
<html>
<head>
	  <title>API Docs</title>
	  <meta charset="utf-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
{{range .}}
	<h1><a href="/{{.Spec.InstanceName | lower}}/doc/swagger/index.html"> {{.Spec.InstanceName}} </a></h1>
	<p>{{.Spec.Description}}</p>
{{end}}
</body>
</html>
`
	tmpl := template.New("index")
	tmpl.Funcs(template.FuncMap{
		"lower": strings.ToLower,
	})
	tmpl, _ = tmpl.Parse(indexHTML)
	var sb strings.Builder
	_ = tmpl.Execute(&sb, targets)
	return sb.String()
}
