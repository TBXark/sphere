package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (a *Auth[ID, USERNAME]) NewAuthMiddleware(abortOnError bool) gin.HandlerFunc {
	abort := func(ctx *gin.Context) {
		if abortOnError {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
		}
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader(AuthorizationHeader)
		if token == "" {
			abort(ctx)
			return
		}

		if len(a.prefix) > 0 && strings.HasPrefix(token, a.prefix+" ") {
			token = token[len(a.prefix)+1:]
		}

		claims, err := a.parser.ParseToken(token)
		if err != nil {
			abort(ctx)
			return
		}

		id, _ := strconv.Atoi(claims.Subject)
		ctx.Set(ContextKeyID, int64(id))
		ctx.Set(ContextKeyUsername, claims.Username)
		ctx.Set(ContextKeyRoles, claims.Username)
	}
}

func (a *Auth[I, U]) NewPermissionMiddleware(resource string, acl AccessControl) gin.HandlerFunc {
	abort := func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
	}
	return func(ctx *gin.Context) {
		rolesRaw, exist := ctx.Get(ContextKeyRoles)
		if !exist {
			abort(ctx)
			return
		}

		roleStr, ok := rolesRaw.(string)
		if !ok {
			abort(ctx)
			return
		}

		roles := a.parser.ParseRoles(roleStr)
		for _, r := range roles {
			if acl.IsAllowed(r, resource) {
				return
			}
		}

		abort(ctx)
	}
}
