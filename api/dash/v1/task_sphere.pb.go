// Code generated by protoc-gen-sphere. DO NOT EDIT.
// versions:
// - protoc             (unknown)
// source: dash/v1/task.proto

package dashv1

import (
	context "context"
	protovalidate_go "github.com/bufbuild/protovalidate-go"
	gin "github.com/gin-gonic/gin"
	ginx "github.com/tbxark/sphere/pkg/server/ginx"
)

var _ = new(context.Context)
var _ = new(gin.Context)
var _ = new(ginx.ErrorResponse)
var _ = new(protovalidate_go.Validator)

const OperationTaskServiceTaskDetail = "/dash.v1.TaskService/TaskDetail"
const OperationTaskServiceTaskList = "/dash.v1.TaskService/TaskList"

var TaskServiceOperationRoutes = [...][3]string{
	{OperationTaskServiceTaskList, "GET", "/api/task/list"},
	{OperationTaskServiceTaskDetail, "GET", "/api/task/detail/:id"},
}

type TaskServiceHTTPServer interface {
	TaskDetail(context.Context, *TaskDetailRequest) (*TaskDetailResponse, error)
	TaskList(context.Context, *TaskListRequest) (*TaskListResponse, error)
}

// @Summary TaskList
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query integer false "page"
// @Param size query integer false "size"
// @Success 200 {object} ginx.DataResponse[TaskListResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/task/list [get]
func _TaskService_TaskList0_HTTP_Handler(srv TaskServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*TaskListResponse, error) {
		var in TaskListRequest
		if err := ginx.ShouldBindQuery(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.TaskList(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

// @Summary TaskDetail
// @Tags dash.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Success 200 {object} ginx.DataResponse[TaskDetailResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/task/detail/{id} [get]
func _TaskService_TaskDetail0_HTTP_Handler(srv TaskServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*TaskDetailResponse, error) {
		var in TaskDetailRequest
		if err := ginx.ShouldBindUri(ctx, &in); err != nil {
			return nil, err
		}
		out, err := srv.TaskDetail(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

func RegisterTaskServiceHTTPServer(route gin.IRouter, srv TaskServiceHTTPServer) {
	r := route.Group("/")
	r.GET("/api/task/list", _TaskService_TaskList0_HTTP_Handler(srv))
	r.GET("/api/task/detail/:id", _TaskService_TaskDetail0_HTTP_Handler(srv))
}
