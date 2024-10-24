// Code generated by protoc-gen-sphere. DO NOT EDIT.
// versions:
// - protoc             (unknown)
// source: dash/v1/auth.proto

package dashv1

import (
	context "context"
	protovalidate_go "github.com/bufbuild/protovalidate-go"
	gin "github.com/gin-gonic/gin"
	ginx "github.com/tbxark/sphere/pkg/server/ginx"
)

var _ = new(context.Context)
var _ = new(gin.Context)
var _ = new(ginx.DataResponse[string])
var _ = new(protovalidate_go.Validator)

type AuthServiceHTTPServer interface {
	AuthLogin(context.Context, *AuthLoginRequest) (*AuthLoginResponse, error)
	AuthRefresh(context.Context, *AuthRefreshRequest) (*AuthRefreshResponse, error)
}

// @Summary AuthLogin
// @Tags dash.v1
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer token"
// @Param request body AuthLoginRequest true "Request body"
// @Success 200 {object} ginx.DataResponse[AuthLoginResponse]
// @Router /api/auth/login [post]
func _AuthService_AuthLogin0_HTTP_Handler(srv AuthServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*AuthLoginResponse, error) {
		var in AuthLoginRequest
		if err := ginx.ShouldBindJSON(ctx, &in); err != nil {
			return nil, err
		}
		if err := protovalidate_go.Validate(&in); err != nil {
			return nil, err
		}
		out, err := srv.AuthLogin(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary AuthRefresh
// @Tags dash.v1
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer token"
// @Param request body AuthRefreshRequest true "Request body"
// @Success 200 {object} ginx.DataResponse[AuthRefreshResponse]
// @Router /api/auth/refresh [post]
func _AuthService_AuthRefresh0_HTTP_Handler(srv AuthServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*AuthRefreshResponse, error) {
		var in AuthRefreshRequest
		if err := ginx.ShouldBindJSON(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.AuthRefresh(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

func RegisterAuthServiceHTTPServer(route gin.IRouter, srv AuthServiceHTTPServer) {
	r := route.Group("/")
	r.POST("/api/auth/login", _AuthService_AuthLogin0_HTTP_Handler(srv))
	r.POST("/api/auth/refresh", _AuthService_AuthRefresh0_HTTP_Handler(srv))
}
