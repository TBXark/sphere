// Code generated by protoc-gen-sphere. DO NOT EDIT.
// versions:
// - protoc             (unknown)
// source: api/v1/system.proto

package apiv1

import (
	context "context"
	ginx "github.com/TBXark/sphere/pkg/server/ginx"
	gin "github.com/gin-gonic/gin"
)

var _ = new(context.Context)
var _ = new(gin.IRouter)
var _ = new(ginx.ErrorResponse)

const OperationSystemServiceStatus = "/api.v1.SystemService/Status"

var EndpointsSystemService = [...][3]string{
	{OperationSystemServiceStatus, "GET", "/api/status"},
}

type SystemServiceHTTPServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
}

// @Summary Status
// @Tags api.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} ginx.DataResponse[StatusResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/status [get]
func _SystemService_Status0_HTTP_Handler(srv SystemServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*StatusResponse, error) {
		var in StatusRequest
		out, err := srv.Status(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

func RegisterSystemServiceHTTPServer(route gin.IRouter, srv SystemServiceHTTPServer) {
	r := route.Group("/")
	r.GET("/api/status", _SystemService_Status0_HTTP_Handler(srv))
}
