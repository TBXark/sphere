// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplateAPI = `{
    "schemes": {{ marshal .Schemes }},
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
                    "api"
                ],
                "summary": "获取系统状态",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_web.DataResponse-github_com_tbxark_go-base-api_pkg_web_models_MessageResponse"
                        }
                    }
                }
            }
        },
        "/api/upload/token": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "获取上传凭证",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件名",
                        "name": "filename",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_web.DataResponse-internal_biz_api_UploadTokenResponse"
                        }
                    }
                }
            }
        },
        "/api/user/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "获取当前用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_web.DataResponse-internal_biz_api_UserInfoMeResponse"
                        }
                    }
                }
            }
        },
        "/api/user/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_biz_api.UpdateUserInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_web.DataResponse-internal_biz_api_UpdateUserInfoResponse"
                        }
                    }
                }
            }
        },
        "/api/wx/mini/auth": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "微信小程序登录",
                "parameters": [
                    {
                        "description": "登录信息",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_biz_api.WxMiniAuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_web.DataResponse-internal_biz_api_AuthResponse"
                        }
                    }
                }
            }
        },
        "/api/wx/mini/bind/phone": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "绑定手机号",
                "parameters": [
                    {
                        "description": "绑定信息",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_biz_api.WxMiniBindPhoneRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_web.DataResponse-github_com_tbxark_go-base-api_pkg_web_models_MessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_tbxark_go-base-api_pkg_dao_ent.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "created_at": {
                    "description": "创建时间",
                    "type": "integer"
                },
                "flags": {
                    "description": "标记位",
                    "type": "integer"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "github_com_tbxark_go-base-api_pkg_web.DataResponse-github_com_tbxark_go-base-api_pkg_web_models_MessageResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_web_models.MessageResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github_com_tbxark_go-base-api_pkg_web.DataResponse-internal_biz_api_AuthResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/internal_biz_api.AuthResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github_com_tbxark_go-base-api_pkg_web.DataResponse-internal_biz_api_UpdateUserInfoResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/internal_biz_api.UpdateUserInfoResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github_com_tbxark_go-base-api_pkg_web.DataResponse-internal_biz_api_UploadTokenResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/internal_biz_api.UploadTokenResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github_com_tbxark_go-base-api_pkg_web.DataResponse-internal_biz_api_UserInfoMeResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/internal_biz_api.UserInfoMeResponse"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github_com_tbxark_go-base-api_pkg_web_models.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "internal_biz_api.AuthResponse": {
            "type": "object",
            "properties": {
                "isNew": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_dao_ent.User"
                }
            }
        },
        "internal_biz_api.UpdateUserInfoRequest": {
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
        "internal_biz_api.UpdateUserInfoResponse": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_dao_ent.User"
                }
            }
        },
        "internal_biz_api.UploadTokenResponse": {
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
        "internal_biz_api.UserInfoMeResponse": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_dao_ent.User"
                },
                "inviter": {
                    "$ref": "#/definitions/github_com_tbxark_go-base-api_pkg_dao_ent.User"
                }
            }
        },
        "internal_biz_api.WxMiniAuthRequest": {
            "type": "object",
            "required": [
                "code"
            ],
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "internal_biz_api.WxMiniBindPhoneRequest": {
            "type": "object",
            "required": [
                "code"
            ],
            "properties": {
                "code": {
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
    }
}`

// SwaggerInfoAPI holds exported Swagger Info so clients can modify it
var SwaggerInfoAPI = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "API",
	SwaggerTemplate:  docTemplateAPI,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfoAPI.InstanceName(), SwaggerInfoAPI)
}
