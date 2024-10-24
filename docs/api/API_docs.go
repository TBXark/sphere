// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplateAPI = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/status": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api.v1"
                ],
                "summary": "Status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-apiv1_StatusResponse"
                        }
                    }
                }
            }
        },
        "/api/test/{path_test1}/second/{path_test2}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shared.v1"
                ],
                "summary": "RunTest",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "path_test1",
                        "name": "path_test1",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "path_test2",
                        "name": "path_test2",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "query_test1",
                        "name": "query_test1",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "query_test2",
                        "name": "query_test2",
                        "in": "query"
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sharedv1.RunTestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-sharedv1_RunTestResponse"
                        }
                    }
                }
            }
        },
        "/api/upload/token": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shared.v1"
                ],
                "summary": "UploadToken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sharedv1.UploadTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-sharedv1_UploadTokenResponse"
                        }
                    }
                }
            }
        },
        "/api/user/bind/phone/wxmini": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api.v1"
                ],
                "summary": "BindPhoneWxMini",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apiv1.BindPhoneWxMiniRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-apiv1_BindPhoneWxMiniResponse"
                        }
                    }
                }
            }
        },
        "/api/user/me": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api.v1"
                ],
                "summary": "Me",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-apiv1_MeResponse"
                        }
                    }
                }
            }
        },
        "/api/user/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api.v1"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apiv1.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-apiv1_UpdateResponse"
                        }
                    }
                }
            }
        },
        "/v1/auth/wxmini": {
            "post": {
                "description": "请求小程序",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api.v1"
                ],
                "summary": "AuthWxMini",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apiv1.AuthWxMiniRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-apiv1_AuthWxMiniResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apiv1.AuthWxMiniRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "apiv1.AuthWxMiniResponse": {
            "type": "object",
            "properties": {
                "is_new": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/sharedv1.User"
                }
            }
        },
        "apiv1.BindPhoneWxMiniRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "apiv1.BindPhoneWxMiniResponse": {
            "type": "object"
        },
        "apiv1.MeResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/sharedv1.User"
                }
            }
        },
        "apiv1.StatusResponse": {
            "type": "object"
        },
        "apiv1.UpdateRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "apiv1.UpdateResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/sharedv1.User"
                }
            }
        },
        "ginx.DataResponse-apiv1_AuthWxMiniResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/apiv1.AuthWxMiniResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-apiv1_BindPhoneWxMiniResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/apiv1.BindPhoneWxMiniResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-apiv1_MeResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/apiv1.MeResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-apiv1_StatusResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/apiv1.StatusResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-apiv1_UpdateResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/apiv1.UpdateResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-sharedv1_RunTestResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/sharedv1.RunTestResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-sharedv1_UploadTokenResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/sharedv1.UploadTokenResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "sharedv1.RunTestRequest": {
            "type": "object",
            "properties": {
                "field_test1": {
                    "type": "string"
                },
                "field_test2": {
                    "type": "integer"
                }
            }
        },
        "sharedv1.RunTestResponse": {
            "type": "object",
            "properties": {
                "field_test1": {
                    "type": "string"
                },
                "field_test2": {
                    "type": "integer"
                },
                "path_test1": {
                    "type": "string"
                },
                "path_test2": {
                    "type": "integer"
                },
                "query_test1": {
                    "type": "string"
                },
                "query_test2": {
                    "type": "integer"
                }
            }
        },
        "sharedv1.UploadTokenRequest": {
            "type": "object",
            "properties": {
                "filename": {
                    "type": "string"
                }
            }
        },
        "sharedv1.UploadTokenResponse": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "sharedv1.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "JWT token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "security": [
        {
            "ApiKeyAuth ": [
                ""
            ]
        }
    ]
}`

// SwaggerInfoAPI holds exported Swagger Info so clients can modify it
var SwaggerInfoAPI = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "sphere",
	Description:      "sphere api docs",
	InfoInstanceName: "API",
	SwaggerTemplate:  docTemplateAPI,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfoAPI.InstanceName(), SwaggerInfoAPI)
}
