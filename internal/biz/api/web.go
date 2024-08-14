package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tbxark/go-base-api/internal/pkg/dao"
	"github.com/tbxark/go-base-api/internal/pkg/render"
	"github.com/tbxark/go-base-api/pkg/cache"
	"github.com/tbxark/go-base-api/pkg/cdn"
	"github.com/tbxark/go-base-api/pkg/log"
	"github.com/tbxark/go-base-api/pkg/log/field"
	"github.com/tbxark/go-base-api/pkg/web/auth/jwt_tokens"
	"github.com/tbxark/go-base-api/pkg/web/middleware/auth"
	"github.com/tbxark/go-base-api/pkg/web/middleware/logger"
	"github.com/tbxark/go-base-api/pkg/wechat"
	"golang.org/x/sync/singleflight"
	"net/http"
	"strconv"
	"strings"
)

type Config struct {
	JWT     string `json:"jwt"`
	Address string `json:"address"`
}

type Web struct {
	config *Config
	Engine *gin.Engine
	sf     singleflight.Group
	db     *dao.Dao
	wx     *wechat.Wechat
	cdn    cdn.CDN
	cache  cache.ByteCache
	render *render.Render
	token  *jwt_tokens.JwtAuth
	auth   *auth.Auth
}

func NewWebServer(config *Config, db *dao.Dao, wx *wechat.Wechat, cdn cdn.CDN, cache cache.ByteCache) *Web {
	token := jwt_tokens.NewJwtAuth(config.JWT)
	return &Web{
		config: config,
		Engine: gin.New(),
		wx:     wx,
		db:     db,
		cdn:    cdn,
		cache:  cache,
		render: render.NewRender(cdn, db, true),
		token:  token,
		auth:   auth.NewJwtAuth(jwt_tokens.AuthorizationPrefixBearer, token),
	}
}

func (w *Web) Identifier() string {
	return "api"
}

func (w *Web) Run() {
	zapLogger := log.ZapLogger().With(field.String("module", "api"))
	loggerMiddleware := logger.NewZapLoggerMiddleware(zapLogger)
	recoveryMiddleware := logger.NewZapRecoveryMiddleware(zapLogger)
	//rateLimiter := middleware.NewNewRateLimiterByClientIP(100*time.Millisecond, 10, time.Hour)

	w.Engine.Use(loggerMiddleware, recoveryMiddleware)

	api := w.Engine.Group("/", w.auth.NewAuthMiddleware(false))

	w.bindAuthRoute(api)
	w.bindUserRoute(api)
	w.bindSystemRoute(api)

	err := w.Engine.Run(w.config.Address)
	if err != nil {
		log.Warnw("api server run error", field.Error(err))
	}
}

func (w *Web) uploadRemoteImage(ctx *gin.Context, url string) (string, error) {
	key := w.cdn.KeyFromURL(url)
	if key == "" {
		return key, nil
	}
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		return key, nil
	}
	id, err := w.auth.GetCurrentID(ctx)
	if err != nil {
		return "", err
	}
	key = cdn.DefaultKeyBuilder(strconv.Itoa(id))(url, "user")
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	ret, err := w.cdn.UploadFile(ctx, resp.Body, resp.ContentLength, key)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}
