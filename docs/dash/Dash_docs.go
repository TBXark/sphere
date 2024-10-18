// Package dash Code generated by swaggo/swag. DO NOT EDIT
package dash

import "github.com/swaggo/swag"

const docTemplateDash = `{
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
        "/api/admin/create": {
            "post": {
                "description": "AdminCreate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dash.v1"
                ],
                "summary": "AdminCreate",
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
                            "$ref": "#/definitions/dashv1.AdminCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-dashv1_AdminCreateResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/delete/{id}": {
            "delete": {
                "description": "AdminDelete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dash.v1"
                ],
                "summary": "AdminDelete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-dashv1_AdminDeleteResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/detail/{id}": {
            "get": {
                "description": "AdminDetail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dash.v1"
                ],
                "summary": "AdminDetail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-dashv1_AdminDetailResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/list": {
            "get": {
                "description": "AdminList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dash.v1"
                ],
                "summary": "AdminList",
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
                            "$ref": "#/definitions/ginx.DataResponse-dashv1_AdminListResponse"
                        }
                    }
                }
            }
        },
        "/api/admin/update/{id}": {
            "post": {
                "description": "AdminUpdate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dash.v1"
                ],
                "summary": "AdminUpdate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dashv1.AdminUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-dashv1_AdminUpdateResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/login": {
            "post": {
                "description": "AuthLogin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dash.v1"
                ],
                "summary": "AuthLogin",
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
                            "$ref": "#/definitions/dashv1.AuthLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-dashv1_AuthLoginResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/refresh": {
            "post": {
                "description": "AuthRefresh",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dash.v1"
                ],
                "summary": "AuthRefresh",
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
                            "$ref": "#/definitions/dashv1.AuthRefreshRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-dashv1_AuthRefreshResponse"
                        }
                    }
                }
            }
        },
        "/api/cache/reset": {
            "post": {
                "description": "CacheReset",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dash.v1"
                ],
                "summary": "CacheReset",
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
                            "$ref": "#/definitions/dashv1.CacheResetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ginx.DataResponse-dashv1_CacheResetResponse"
                        }
                    }
                }
            }
        },
        "/api/test/{path_test1}/second/{path_test2}": {
            "post": {
                "description": "RunTest",
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
                "description": "UploadToken",
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
        }
    },
    "definitions": {
        "dashv1.Admin": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dashv1.AdminCreateRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dashv1.AdminCreateResponse": {
            "type": "object",
            "properties": {
                "admin": {
                    "$ref": "#/definitions/dashv1.Admin"
                }
            }
        },
        "dashv1.AdminDeleteResponse": {
            "type": "object"
        },
        "dashv1.AdminDetailResponse": {
            "type": "object",
            "properties": {
                "admin": {
                    "$ref": "#/definitions/dashv1.Admin"
                }
            }
        },
        "dashv1.AdminListResponse": {
            "type": "object",
            "properties": {
                "admins": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dashv1.Admin"
                    }
                }
            }
        },
        "dashv1.AdminUpdateRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dashv1.AdminUpdateResponse": {
            "type": "object",
            "properties": {
                "admin": {
                    "$ref": "#/definitions/dashv1.Admin"
                }
            }
        },
        "dashv1.AuthLoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dashv1.AuthLoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "expires": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dashv1.AuthRefreshRequest": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dashv1.AuthRefreshResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dashv1.CacheResetRequest": {
            "type": "object"
        },
        "dashv1.CacheResetResponse": {
            "type": "object"
        },
        "ginx.DataResponse-dashv1_AdminCreateResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dashv1.AdminCreateResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-dashv1_AdminDeleteResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dashv1.AdminDeleteResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-dashv1_AdminDetailResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dashv1.AdminDetailResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-dashv1_AdminListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dashv1.AdminListResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-dashv1_AdminUpdateResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dashv1.AdminUpdateResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-dashv1_AuthLoginResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dashv1.AuthLoginResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-dashv1_AuthRefreshResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dashv1.AuthRefreshResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "ginx.DataResponse-dashv1_CacheResetResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dashv1.CacheResetResponse"
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

// SwaggerInfoDash holds exported Swagger Info so clients can modify it
var SwaggerInfoDash = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "sphere",
	Description:      "sphere api docs",
	InfoInstanceName: "Dash",
	SwaggerTemplate:  docTemplateDash,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfoDash.InstanceName(), SwaggerInfoDash)
}
