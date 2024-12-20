// Code generated by protoc-gen-sphere. DO NOT EDIT.
// versions:
// - protoc             (unknown)
// source: dash/v1/system.proto

package dashv1

import (
	context "context"
	ginx "github.com/TBXark/sphere/pkg/server/ginx"
	gin "github.com/gin-gonic/gin"
)

var _ = new(context.Context)
var _ = new(gin.IRouter)
var _ = new(ginx.ErrorResponse)

const OperationSystemServiceCacheReset = "/dash.v1.SystemService/CacheReset"
const OperationSystemServiceMenuAll = "/dash.v1.SystemService/MenuAll"

var EndpointsSystemService = [...][3]string{
	{OperationSystemServiceCacheReset, "POST", "/api/cache/reset"},
	{OperationSystemServiceMenuAll, "GET", "/api/menu/all"},
}

type SystemServiceHTTPServer interface {
	CacheReset(context.Context, *CacheResetRequest) (*CacheResetResponse, error)
	MenuAll(context.Context, *MenuAllRequest) (*MenuAllResponse, error)
}

// @Summary CacheReset
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body CacheResetRequest true "Request body"
// @Success 200 {object} ginx.DataResponse[CacheResetResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/cache/reset [post]
func _SystemService_CacheReset0_HTTP_Handler(srv SystemServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*CacheResetResponse, error) {
		var in CacheResetRequest
		if err := ginx.ShouldBindJSON(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.CacheReset(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary MenuAll
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} ginx.DataResponse[MenuAllResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/menu/all [get]
func _SystemService_MenuAll0_HTTP_Handler(srv SystemServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*MenuAllResponse, error) {
		var in MenuAllRequest
		out, err := srv.MenuAll(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

func RegisterSystemServiceHTTPServer(route gin.IRouter, srv SystemServiceHTTPServer) {
	r := route.Group("/")
	r.POST("/api/cache/reset", _SystemService_CacheReset0_HTTP_Handler(srv))
	r.GET("/api/menu/all", _SystemService_MenuAll0_HTTP_Handler(srv))
}
