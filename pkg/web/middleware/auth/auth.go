package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tbxark/go-base-api/pkg/web/auth/parser"
	"net/http"
	"strconv"
	"strings"
)

type Auth struct {
	*Base
	prefix string
	parser parser.AuthParser
}

func NewAuth(prefix string, parser parser.AuthParser) *Auth {
	return &Auth{
		Base:   &Base{},
		prefix: prefix,
		parser: parser,
	}
}

func (w *Auth) parseToken(token string) (claims *parser.Claims, err error) {
	if len(w.prefix) > 0 && strings.HasPrefix(token, w.prefix+" ") {
		token = token[len(w.prefix)+1:]
	}
	return w.parser.ParseToken(token)
}

func (w *Auth) NewAuthMiddleware(abortOnError bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(AuthorizationHeader)
		if token == "" {
			w.handleUnauthorized(ctx, abortOnError)
			return
		}

		claims, err := w.parseToken(token)
		if err != nil {
			w.handleUnauthorized(ctx, abortOnError)
			return
		}
		w.setContextValues(ctx, claims)
	}
}

func (w *Auth) handleUnauthorized(ctx *gin.Context, abortOnError bool) {
	if abortOnError {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}
}

func (w *Auth) setContextValues(ctx *gin.Context, claims *parser.Claims) {
	id, _ := strconv.Atoi(claims.Subject)
	ctx.Set(ContextKeyID, id)
	ctx.Set(ContextKeyUsername, claims.Username)
	ctx.Set(ContextKeyRoles, claims.Username)
}

func (w *Auth) NewPermissionMiddleware(resource string, acl *ACL) func(context *gin.Context) {
	return func(ctx *gin.Context) {
		rolesRaw, exist := ctx.Get(ContextKeyRoles)
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
			})
			return
		}
		roleStr, ok := rolesRaw.(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
			})
			return
		}
		roles := w.parser.ParseRoles(roleStr)
		for r := range roles {
			if acl.IsAllowed(r, resource) {
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
	}
}
