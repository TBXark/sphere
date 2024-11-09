// Code generated by protoc-gen-sphere. DO NOT EDIT.
// versions:
// - protoc             (unknown)
// source: dash/v1/admin.proto

package dashv1

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

const OperationAdminServiceAdminCreate = "/dash.v1.AdminService/AdminCreate"
const OperationAdminServiceAdminDelete = "/dash.v1.AdminService/AdminDelete"
const OperationAdminServiceAdminDetail = "/dash.v1.AdminService/AdminDetail"
const OperationAdminServiceAdminList = "/dash.v1.AdminService/AdminList"
const OperationAdminServiceAdminRoleList = "/dash.v1.AdminService/AdminRoleList"
const OperationAdminServiceAdminUpdate = "/dash.v1.AdminService/AdminUpdate"

var EndpointsAdminService = [...][3]string{
	{OperationAdminServiceAdminList, "GET", "/api/admin/list"},
	{OperationAdminServiceAdminCreate, "POST", "/api/admin/create"},
	{OperationAdminServiceAdminUpdate, "POST", "/api/admin/update/:id"},
	{OperationAdminServiceAdminDetail, "GET", "/api/admin/detail/:id"},
	{OperationAdminServiceAdminDelete, "DELETE", "/api/admin/delete/:id"},
	{OperationAdminServiceAdminRoleList, "GET", "/api/admin/role/list"},
}

type AdminServiceHTTPServer interface {
	AdminCreate(context.Context, *AdminCreateRequest) (*AdminCreateResponse, error)
	AdminDelete(context.Context, *AdminDeleteRequest) (*AdminDeleteResponse, error)
	AdminDetail(context.Context, *AdminDetailRequest) (*AdminDetailResponse, error)
	AdminList(context.Context, *AdminListRequest) (*AdminListResponse, error)
	AdminRoleList(context.Context, *AdminRoleListRequest) (*AdminRoleListResponse, error)
	AdminUpdate(context.Context, *AdminUpdateRequest) (*AdminUpdateResponse, error)
}

// @Summary AdminList
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} ginx.DataResponse[AdminListResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/admin/list [get]
func _AdminService_AdminList0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*AdminListResponse, error) {
		var in AdminListRequest
		out, err := srv.AdminList(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary AdminCreate
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body AdminCreateRequest true "Request body"
// @Success 200 {object} ginx.DataResponse[AdminCreateResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/admin/create [post]
func _AdminService_AdminCreate0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*AdminCreateResponse, error) {
		var in AdminCreateRequest
		if err := ginx.ShouldBindJSON(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.AdminCreate(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary AdminUpdate
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Param request body AdminUpdateRequest true "Request body"
// @Success 200 {object} ginx.DataResponse[AdminUpdateResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/admin/update/{id} [post]
func _AdminService_AdminUpdate0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*AdminUpdateResponse, error) {
		var in AdminUpdateRequest
		if err := ginx.ShouldBindJSON(ctx, &in); err != nil {
			return nil, err
		}
		if err := ginx.ShouldBindUri(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.AdminUpdate(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary AdminDetail
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Success 200 {object} ginx.DataResponse[AdminDetailResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/admin/detail/{id} [get]
func _AdminService_AdminDetail0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*AdminDetailResponse, error) {
		var in AdminDetailRequest
		if err := ginx.ShouldBindUri(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.AdminDetail(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary AdminDelete
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Success 200 {object} ginx.DataResponse[AdminDeleteResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/admin/delete/{id} [delete]
func _AdminService_AdminDelete0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*AdminDeleteResponse, error) {
		var in AdminDeleteRequest
		if err := ginx.ShouldBindUri(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.AdminDelete(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary AdminRoleList
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} ginx.DataResponse[AdminRoleListResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/admin/role/list [get]
func _AdminService_AdminRoleList0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*AdminRoleListResponse, error) {
		var in AdminRoleListRequest
		out, err := srv.AdminRoleList(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

func RegisterAdminServiceHTTPServer(route gin.IRouter, srv AdminServiceHTTPServer) {
	r := route.Group("/")
	r.GET("/api/admin/list", _AdminService_AdminList0_HTTP_Handler(srv))
	r.POST("/api/admin/create", _AdminService_AdminCreate0_HTTP_Handler(srv))
	r.POST("/api/admin/update/:id", _AdminService_AdminUpdate0_HTTP_Handler(srv))
	r.GET("/api/admin/detail/:id", _AdminService_AdminDetail0_HTTP_Handler(srv))
	r.DELETE("/api/admin/delete/:id", _AdminService_AdminDelete0_HTTP_Handler(srv))
	r.GET("/api/admin/role/list", _AdminService_AdminRoleList0_HTTP_Handler(srv))
}
