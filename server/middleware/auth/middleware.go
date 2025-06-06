package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/TBXark/sphere/server/auth/authorizer"
	"github.com/gin-gonic/gin"
)

var errTokenNotFound = errors.New("token not found")

const (
	AuthorizationHeader = "Authorization"
)

type AccessControl interface {
	IsAllowed(role, resource string) bool
}

func abortUnauthorized(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
}

func abortForbidden(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": "forbidden",
	})
}

func parserToken[T authorizer.UID](ctx *gin.Context, token string, transform func(text string) (string, error), parser authorizer.Parser[authorizer.RBACClaims[T]]) error {
	if token == "" {
		return errTokenNotFound
	}
	if transform != nil {
		tranToken, err := transform(token)
		if err != nil {
			return err
		}
		if tranToken == "" {
			return errTokenNotFound
		}
		token = tranToken
	}
	claims, err := parser.ParseToken(ctx, token)
	if err != nil {
		return err
	}
	ctx.Set(authorizer.ContextKeyUID, claims.UID)
	ctx.Set(authorizer.ContextKeySubject, claims.Subject)
	ctx.Set(authorizer.ContextKeyRoles, claims.Roles)
	return nil
}

func NewAuthMiddleware[T authorizer.UID](prefix string, parser authorizer.Parser[authorizer.RBACClaims[T]], abortOnError bool) gin.HandlerFunc {
	prefix = strings.TrimSpace(prefix)
	if len(prefix) > 0 {
		prefix = prefix + " "
	}
	return NewCommonAuthMiddleware[T](
		func(ctx *gin.Context) (string, error) {
			return ctx.GetHeader(AuthorizationHeader), nil
		},
		func(token string) (string, error) {
			if len(prefix) > 0 && strings.HasPrefix(token, prefix) {
				token = strings.TrimSpace(strings.TrimPrefix(token, prefix))
			}
			return token, nil
		},
		parser,
		abortOnError,
	)
}

func NewCookieAuthMiddleware[T authorizer.UID](cookieName string, transform func(raw string) (string, error), parser authorizer.Parser[authorizer.RBACClaims[T]], abortOnError bool) gin.HandlerFunc {
	return NewCommonAuthMiddleware[T](func(ctx *gin.Context) (string, error) {
		return ctx.Cookie(cookieName)
	}, transform, parser, abortOnError)
}

func NewCommonAuthMiddleware[T authorizer.UID](loader func(ctx *gin.Context) (string, error), transform func(text string) (string, error), parser authorizer.Parser[authorizer.RBACClaims[T]], abortOnError bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := loader(ctx)
		if err != nil && abortOnError {
			abortUnauthorized(ctx)
			return
		}
		err = parserToken(ctx, token, transform, parser)
		if err != nil && abortOnError {
			abortUnauthorized(ctx)
			return
		}
		ctx.Next()
	}
}

func NewPermissionMiddleware(resource string, acl AccessControl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isAllowed := false
		if rolesRaw, exist := ctx.Get(authorizer.ContextKeyRoles); exist {
			if roles, ok := rolesRaw.([]string); ok {
				for _, r := range roles {
					if acl.IsAllowed(r, resource) {
						isAllowed = true
						break
					}
				}
			}
		}
		if isAllowed {
			ctx.Next()
		} else {
			abortForbidden(ctx)
			return
		}
	}
}
