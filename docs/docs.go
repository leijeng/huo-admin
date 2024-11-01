// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "lei",
            "url": "https://github.com/leijeng/huo-admin",
            "email": "lisite199505@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v2/admin/sys/sys-role/create": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "sys-SysRole"
                ],
                "summary": "创建SysRole",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "团队id",
                        "name": "teamId",
                        "in": "header"
                    },
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SysRoleDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SysRole"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v2/admin/sys/sys-role/del": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "sys-SysRole"
                ],
                "summary": "删除SysRole",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "团队id",
                        "name": "teamId",
                        "in": "header"
                    },
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/base.ReqIds"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SysRole"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v2/admin/sys/sys-role/get": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "sys-SysRole"
                ],
                "summary": "获取SysRole",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "团队id",
                        "name": "teamId",
                        "in": "header"
                    },
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/base.ReqId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SysRole"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v2/admin/sys/sys-role/page": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "sys-SysRole"
                ],
                "summary": "获取SysRole列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "团队id",
                        "name": "teamId",
                        "in": "header"
                    },
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SysRoleGetPageReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/base.PageResp"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.SysRole"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v2/admin/sys/sys-role/update": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "sys-SysRole"
                ],
                "summary": "更新SysRole",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "团队id",
                        "name": "teamId",
                        "in": "header"
                    },
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SysRoleDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SysRole"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.PageResp": {
            "type": "object",
            "properties": {
                "currentPage": {
                    "description": "当前第几页",
                    "type": "integer"
                },
                "list": {
                    "description": "数据列表"
                },
                "pageSize": {
                    "description": "分页大小",
                    "type": "integer"
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                }
            }
        },
        "base.ReqId": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "base.ReqIds": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "多id",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "base.Resp": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "返回码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "msg": {
                    "description": "消息",
                    "type": "string"
                },
                "reqId": {
                    "description": "` + "`" + `json:\"请求id\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "dto.SysRoleDto": {
            "type": "object",
            "properties": {
                "createdBy": {
                    "description": "创建人id",
                    "type": "integer"
                },
                "id": {
                    "description": "主键",
                    "type": "integer"
                },
                "name": {
                    "description": "角色名称",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "roleKey": {
                    "description": "角色代码",
                    "type": "string"
                },
                "roleSort": {
                    "description": "排序",
                    "type": "integer"
                },
                "status": {
                    "description": "状态：1 有效 2 无效",
                    "type": "integer"
                },
                "teamId": {
                    "description": "团队id",
                    "type": "integer"
                },
                "updatedBy": {
                    "description": "修改人id",
                    "type": "integer"
                }
            }
        },
        "dto.SysRoleGetPageReq": {
            "type": "object",
            "properties": {
                "page": {
                    "description": "页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页大小",
                    "type": "integer"
                },
                "status": {
                    "description": "状态：1 有效 2 无效",
                    "type": "integer"
                }
            }
        },
        "models.SysRole": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "description": "创建人id",
                    "type": "integer"
                },
                "id": {
                    "description": "主键",
                    "type": "integer"
                },
                "name": {
                    "description": "角色名称",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "roleKey": {
                    "description": "角色代码",
                    "type": "string"
                },
                "roleSort": {
                    "description": "排序",
                    "type": "integer"
                },
                "status": {
                    "description": "状态：1 有效 2 无效",
                    "type": "integer"
                },
                "teamId": {
                    "description": "团队id",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "updatedBy": {
                    "description": "修改人id",
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "V1.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Huo API",
	Description:      "致力于做一个开发快速，运行稳定的框架",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
