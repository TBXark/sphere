// Code generated by protoc-gen-sphere. DO NOT EDIT.
// versions:
// - protoc             (unknown)
// source: shared/v1/test.proto

package sharedv1

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

const OperationTestServiceRunTest = "/shared.v1.TestService/RunTest"

type TestServiceHTTPServer interface {
	RunTest(context.Context, *RunTestRequest) (*RunTestResponse, error)
}

// @Summary RunTest
// @Tags shared.v1
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param path_test1 path string true "path_test1"
// @Param path_test2 path integer true "path_test2"
// @Param query_test1 query string false "query_test1"
// @Param query_test2 query integer false "query_test2"
// @Param request body RunTestRequest true "Request body"
// @Success 200 {object} ginx.DataResponse[RunTestResponse]
// @Success 400 {object} ginx.ErrorResponse
// @Success 401 {object} ginx.ErrorResponse
// @Success 403 {object} ginx.ErrorResponse
// @Success 500 {object} ginx.ErrorResponse
// @Router /api/test/{path_test1}/second/{path_test2} [post]
func _TestService_RunTest0_HTTP_Handler(srv TestServiceHTTPServer) func(ctx *gin.Context) {
	return ginx.WithJson(func(ctx *gin.Context) (*RunTestResponse, error) {
		var in RunTestRequest
		if err := ginx.ShouldBindJSON(ctx, &in); err != nil {
			return nil, err
		}
		if err := ginx.ShouldBindQuery(ctx, &in); err != nil {
			return nil, err
		}
		if err := ginx.ShouldBindUri(ctx, &in); err != nil {
			return nil, err
		}
		ctx.Set("operation", OperationTestServiceRunTest)
		out, err := srv.RunTest(ctx, &in)
		if err != nil {
			return nil, err
		}
		return out, nil
	})
}

func RegisterTestServiceHTTPServer(route gin.IRouter, srv TestServiceHTTPServer) {
	r := route.Group("/")
	r.POST("/api/test/:path_test1/second/:path_test2", _TestService_RunTest0_HTTP_Handler(srv))
}

func CreateTestServiceOperationRoute(base string) map[string][]string {
	return map[string][]string{
		OperationTestServiceRunTest: {"POST", ginx.JoinPaths(base, "/api/test/:path_test1/second/:path_test2")},
	}
}
