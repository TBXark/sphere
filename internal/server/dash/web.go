package dash

import (
	"context"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	dashv1 "github.com/tbxark/sphere/api/dash/v1"
	sharedv1 "github.com/tbxark/sphere/api/shared/v1"
	"github.com/tbxark/sphere/internal/service/dash"
	"github.com/tbxark/sphere/internal/service/shared"
	"github.com/tbxark/sphere/pkg/log"
	"github.com/tbxark/sphere/pkg/log/logfields"
	"github.com/tbxark/sphere/pkg/server/auth/acl"
	"github.com/tbxark/sphere/pkg/server/auth/authorizer"
	"github.com/tbxark/sphere/pkg/server/auth/jwtauth"
	"github.com/tbxark/sphere/pkg/server/ginx"
	"github.com/tbxark/sphere/pkg/server/middleware/auth"
	"github.com/tbxark/sphere/pkg/server/middleware/logger"
	"github.com/tbxark/sphere/pkg/server/middleware/ratelimiter"
	"github.com/tbxark/sphere/pkg/server/middleware/selector"
	"github.com/tbxark/sphere/pkg/server/route/cors"
	"github.com/tbxark/sphere/pkg/server/route/pprof"
	"net/http"
	"time"
)

type Web struct {
	config  *Config
	server  *http.Server
	service *dash.Service
}

func NewWebServer(config *Config, service *dash.Service) *Web {
	return &Web{
		config:  config,
		service: service,
	}
}

func (w *Web) Identifier() string {
	return "dash"
}

func (w *Web) Run() error {

	jwtAuthorizer := jwtauth.NewJwtAuth[authorizer.RBACClaims[int64]](w.config.AuthJWT)
	jwtRefresher := jwtauth.NewJwtAuth[authorizer.RBACClaims[int64]](w.config.RefreshJWT)

	zapLogger := log.ZapLogger().With(logfields.String("module", "dash"))
	loggerMiddleware := logger.NewZapLoggerMiddleware(zapLogger)
	recoveryMiddleware := logger.NewZapRecoveryMiddleware(zapLogger)
	authMiddleware := auth.NewAuthMiddleware(jwtauth.AuthorizationPrefixBearer, jwtAuthorizer, true)
	rateLimiter := ratelimiter.NewNewRateLimiterByClientIP(100*time.Millisecond, 10, time.Hour)

	engine := gin.New()
	engine.Use(loggerMiddleware, recoveryMiddleware)

	// 1. 使用go直接反代
	// ignore embed: web.Fs(w.config.DashStatic, nil, "")
	// 2. 使用embed集成
	// with embed: web.Fs("", &dash.Assets, dash.AssetsPath)
	if dashFs, err := ginx.Fs(w.config.HTTP.Static, nil, ""); err == nil && dashFs != nil {
		d := engine.Group("/dash", gzip.Gzip(gzip.DefaultCompression))
		d.StaticFS("/", dashFs)
	}
	// 3. 使用其他服务反代但是允许其跨域访问
	// 其中w.config.DashCors是一个配置项，用于配置允许跨域访问的域名,例如：https://dash.example.com
	if len(w.config.HTTP.Cors) > 0 {
		cors.Setup(engine, w.config.HTTP.Cors)
	}

	api := engine.Group("/")
	needAuthRoute := api.Group("/", authMiddleware)

	w.service.Init(jwtAuthorizer, jwtRefresher)

	if w.config.HTTP.PProf {
		pprof.SetupPProf(api)
	}
	initDefaultRolesACL(w.service.ACL)

	sharedSrc := shared.NewService(w.service.Storage, "dash")
	sharedv1.RegisterStorageServiceHTTPServer(needAuthRoute, sharedSrc)
	sharedv1.RegisterTestServiceHTTPServer(api, sharedSrc)

	authRoute := api.Group("/")
	// 根据元数据限定中间件作用范围
	authRoute.Use(
		selector.NewSelectorMiddleware(
			selector.MatchFunc(
				ginx.MatchOperation(
					authRoute.BasePath(),
					dashv1.EndpointsAuthService[:],
					dashv1.OperationAuthServiceAuthLogin,
				),
			),
			rateLimiter,
		),
	)
	dashv1.RegisterAuthServiceHTTPServer(authRoute, w.service)

	adminRoute := needAuthRoute.Group("/", w.withPermission(dash.PermissionAdmin))
	dashv1.RegisterAdminServiceHTTPServer(adminRoute, w.service)

	systemRoute := needAuthRoute.Group("/")
	dashv1.RegisterSystemServiceHTTPServer(systemRoute, w.service)

	w.server = &http.Server{
		Addr:    w.config.HTTP.Address,
		Handler: engine.Handler(),
	}
	return w.server.ListenAndServe()
}

func (w *Web) Close(ctx context.Context) error {
	if w.server != nil {
		err := w.server.Close()
		if err != nil {
			return err
		}
		w.server = nil
	}
	return nil
}

func (w *Web) withPermission(resource string) gin.HandlerFunc {
	return auth.NewPermissionMiddleware(resource, w.service.ACL)
}

func initDefaultRolesACL(acl *acl.ACL) {
	roles := []string{
		dash.PermissionAdmin,
	}
	for _, r := range roles {
		acl.Allow(dash.PermissionAll, r)
		acl.Allow(r, r)
	}
}
