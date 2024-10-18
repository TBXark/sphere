// Code generated by protoc-gen-sphere. DO NOT EDIT.
// versions:
// - protoc             (unknown)
// source: api/v1/auth.proto

package apiv1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	ginx "github.com/tbxark/sphere/pkg/web/ginx"
)

var _ = new(context.Context)
var _ = new(gin.Context)
var _ = new(ginx.DataResponse[string])

type AuthServiceHTTPServer interface {
	AuthWxMini(context.Context, *AuthWxMiniRequest) (*AuthWxMiniResponse, error)
}

// @Summary AuthWxMini
// @Description AuthWxMini
// @Tags api.v1
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer token"
// @Param request body AuthWxMiniRequest true "Request body"
// @Success 200 {object} ginx.DataResponse[AuthWxMiniResponse]
// @Router /v1/users/create [post]
func _AuthService_AuthWxMini0_HTTP_Handler(srv AuthServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*AuthWxMiniResponse, error) {
		var in AuthWxMiniRequest
		if err := ginx.ShouldBindJSON(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.AuthWxMini(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

func RegisterAuthServiceHTTPServer(route gin.IRouter, srv AuthServiceHTTPServer) {
	r := route.Group("/")
	r.POST("/v1/users/create", _AuthService_AuthWxMini0_HTTP_Handler(srv))
}
