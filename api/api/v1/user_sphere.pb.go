// Code generated by protoc-gen-sphere. DO NOT EDIT.
// versions:
// - protoc             (unknown)
// source: api/v1/user.proto

package apiv1

import (
	context "context"
	protovalidate_go "github.com/bufbuild/protovalidate-go"
	gin "github.com/gin-gonic/gin"
	ginx "github.com/tbxark/sphere/pkg/server/ginx"
)

var _ = new(context.Context)
var _ = new(protovalidate_go.Validator)
var _ = new(gin.IRouter)
var _ = new(ginx.ErrorResponse)

const OperationUserServiceBindPhoneWxMini = "/api.v1.UserService/BindPhoneWxMini"
const OperationUserServiceMe = "/api.v1.UserService/Me"
const OperationUserServiceUpdate = "/api.v1.UserService/Update"

var EndpointsUserService = [...][3]string{
	{OperationUserServiceMe, "GET", "/api/user/me"},
	{OperationUserServiceUpdate, "POST", "/api/user/update"},
	{OperationUserServiceBindPhoneWxMini, "POST", "/api/user/bind/phone/wxmini"},
}

type UserServiceHTTPServer interface {
	BindPhoneWxMini(context.Context, *BindPhoneWxMiniRequest) (*BindPhoneWxMiniResponse, error)
	Me(context.Context, *MeRequest) (*MeResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

// @Summary Me
// @Tags api.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} ginx.DataResponse[MeResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/user/me [get]
func _UserService_Me0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*MeResponse, error) {
		var in MeRequest
		out, err := srv.Me(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary Update
// @Tags api.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body UpdateRequest true "Request body"
// @Success 200 {object} ginx.DataResponse[UpdateResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/user/update [post]
func _UserService_Update0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*UpdateResponse, error) {
		var in UpdateRequest
		if err := ginx.ShouldBindJSON(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.Update(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary BindPhoneWxMini
// @Tags api.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body BindPhoneWxMiniRequest true "Request body"
// @Success 200 {object} ginx.DataResponse[BindPhoneWxMiniResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/user/bind/phone/wxmini [post]
func _UserService_BindPhoneWxMini0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*BindPhoneWxMiniResponse, error) {
		var in BindPhoneWxMiniRequest
		if err := ginx.ShouldBindJSON(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.BindPhoneWxMini(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

func RegisterUserServiceHTTPServer(route gin.IRouter, srv UserServiceHTTPServer) {
	r := route.Group("/")
	r.GET("/api/user/me", _UserService_Me0_HTTP_Handler(srv))
	r.POST("/api/user/update", _UserService_Update0_HTTP_Handler(srv))
	r.POST("/api/user/bind/phone/wxmini", _UserService_BindPhoneWxMini0_HTTP_Handler(srv))
}
