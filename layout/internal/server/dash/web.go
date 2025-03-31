package dash

import (
	"context"
	"github.com/TBXark/sphere/storage"
	"net/http"
	"time"

	dashv1 "github.com/TBXark/sphere/layout/api/dash/v1"
	sharedv1 "github.com/TBXark/sphere/layout/api/shared/v1"
	"github.com/TBXark/sphere/layout/internal/service/dash"
	"github.com/TBXark/sphere/layout/internal/service/shared"
	"github.com/TBXark/sphere/log"
	"github.com/TBXark/sphere/log/logfields"
	"github.com/TBXark/sphere/server/auth/acl"
	"github.com/TBXark/sphere/server/auth/authorizer"
	"github.com/TBXark/sphere/server/auth/jwtauth"
	"github.com/TBXark/sphere/server/ginx"
	"github.com/TBXark/sphere/server/middleware/auth"
	"github.com/TBXark/sphere/server/middleware/logger"
	"github.com/TBXark/sphere/server/middleware/ratelimiter"
	"github.com/TBXark/sphere/server/middleware/selector"
	"github.com/TBXark/sphere/server/route/cors"
	"github.com/TBXark/sphere/server/route/pprof"
	"github.com/gin-gonic/gin"
)

type Web struct {
	config    *Config
	acl       *acl.ACL
	server    *http.Server
	storage   storage.ImageStorage
	service   *dash.Service
	sharedSvc *shared.Service
}

func NewWebServer(config *Config, storage storage.ImageStorage, service *dash.Service) *Web {
	return &Web{
		config:    config,
		service:   service,
		sharedSvc: shared.NewService(storage, "dash"),
	}
}

func (w *Web) Identifier() string {
	return "dash"
}

func (w *Web) Start(ctx context.Context) error {
	jwtAuthorizer := jwtauth.NewJwtAuth[authorizer.RBACClaims[int64]](w.config.AuthJWT)
	jwtRefresher := jwtauth.NewJwtAuth[authorizer.RBACClaims[int64]](w.config.RefreshJWT)

	zapLogger := log.ZapLogger().With(logfields.String("module", "dash"))
	loggerMiddleware := logger.NewZapLoggerMiddleware(zapLogger)
	recoveryMiddleware := logger.NewZapRecoveryMiddleware(zapLogger)
	authMiddleware := auth.NewAuthMiddleware(jwtauth.AuthorizationPrefixBearer, jwtAuthorizer, true)
	rateLimiter := ratelimiter.NewNewRateLimiterByClientIP(100*time.Millisecond, 10, time.Hour)

	engine := gin.New()
	engine.Use(loggerMiddleware, recoveryMiddleware)

	// dashboard 静态资源
	// 1. 不设置 `embed_dash` 编译选项，使用默认的静态资源, 在配置中设置静态资源的绝对路径
	// 2. 设置 `embed_dash` 编译选项，使用内置的静态资源, 静态资源位置在 `assets/dash/dashboard` 目录下
	// 3. 由使用其他服务反代，设置API允许其跨域访问, 其中w.config.DashCors是一个配置项，用于配置允许跨域访问的域名,例如：https://dash.example.com
	w.RegisterDashStatic(engine.Group("/dash"))

	api := engine.Group("/")
	needAuthRoute := api.Group("/", authMiddleware)
	w.service.Init(jwtAuthorizer, jwtRefresher)

	if w.config.HTTP.PProf {
		pprof.SetupPProf(api)
	}
	if len(w.config.HTTP.Cors) > 0 {
		cors.Setup(engine, w.config.HTTP.Cors)
	}
	initDefaultRolesACL(w.acl)

	sharedv1.RegisterStorageServiceHTTPServer(needAuthRoute, w.sharedSvc)
	sharedv1.RegisterTestServiceHTTPServer(api, w.sharedSvc)

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
	authRoute.GET("/api/get-async-routes", ginx.WithJson(func(ctx *gin.Context) ([]struct{}, error) {
		return []struct{}{}, nil
	}))
	dashv1.RegisterAuthServiceHTTPServer(authRoute, w.service)

	adminRoute := needAuthRoute.Group("/", w.withPermission(dash.PermissionAdmin))
	dashv1.RegisterAdminServiceHTTPServer(adminRoute, w.service)

	systemRoute := needAuthRoute.Group("/")
	dashv1.RegisterSystemServiceHTTPServer(systemRoute, w.service)

	w.server = &http.Server{
		Addr:    w.config.HTTP.Address,
		Handler: engine.Handler(),
	}
	return ginx.Start(ctx, w.server, 30*time.Second)
}

func (w *Web) Stop(ctx context.Context) error {
	return ginx.Close(ctx, w.server)
}

func (w *Web) withPermission(resource string) gin.HandlerFunc {
	return auth.NewPermissionMiddleware(resource, w.acl)
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
