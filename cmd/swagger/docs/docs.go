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
            "name": "飞宏IM技术支持",
            "url": "https://github.com/liony823/chat",
            "email": "ilovecoding@foxmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/statistics/group/active": {
            "post": {
                "description": "获取活跃群",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "statistics"
                ],
                "summary": "获取活跃群",
                "operationId": "GetActiveGroup",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/msg.GetActiveGroupReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apiresp.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/msg.GetActiveGroupResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/statistics/group/create": {
            "post": {
                "description": "获取创建群数量",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "statistics"
                ],
                "summary": "获取创建群数量",
                "operationId": "GroupCreateCount",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/group.GroupCreateCountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apiresp.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/group.GroupCreateCountResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/statistics/user/active": {
            "post": {
                "description": "获取活跃用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "statistics"
                ],
                "summary": "获取活跃用户",
                "operationId": "GetActiveUser",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/msg.GetActiveUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apiresp.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/msg.GetActiveUserResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/statistics/user/register": {
            "post": {
                "description": "获取注册用户数",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "statistics"
                ],
                "summary": "获取注册用户数",
                "operationId": "UserRegisterCount",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserRegisterCountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apiresp.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/user.UserRegisterCountResp"
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
        "apiresp.ApiResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "errCode": {
                    "type": "integer"
                },
                "errDlt": {
                    "type": "string"
                },
                "errMsg": {
                    "type": "string"
                }
            }
        },
        "group.GroupCreateCountReq": {
            "type": "object",
            "properties": {
                "end": {
                    "type": "integer"
                },
                "start": {
                    "type": "integer"
                }
            }
        },
        "group.GroupCreateCountResp": {
            "type": "object",
            "properties": {
                "before": {
                    "type": "integer"
                },
                "count": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "msg.ActiveGroup": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "group": {
                    "$ref": "#/definitions/sdkws.GroupInfo"
                }
            }
        },
        "msg.ActiveUser": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/sdkws.UserInfo"
                }
            }
        },
        "msg.GetActiveGroupReq": {
            "type": "object",
            "properties": {
                "ase": {
                    "type": "boolean"
                },
                "end": {
                    "type": "integer"
                },
                "pagination": {
                    "$ref": "#/definitions/sdkws.RequestPagination"
                },
                "start": {
                    "type": "integer"
                }
            }
        },
        "msg.GetActiveGroupResp": {
            "type": "object",
            "properties": {
                "dateCount": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "groupCount": {
                    "type": "integer"
                },
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/msg.ActiveGroup"
                    }
                },
                "msgCount": {
                    "type": "integer"
                }
            }
        },
        "msg.GetActiveUserReq": {
            "type": "object",
            "properties": {
                "ase": {
                    "type": "boolean"
                },
                "end": {
                    "type": "integer"
                },
                "group": {
                    "type": "boolean"
                },
                "pagination": {
                    "$ref": "#/definitions/sdkws.RequestPagination"
                },
                "start": {
                    "type": "integer"
                }
            }
        },
        "msg.GetActiveUserResp": {
            "type": "object",
            "properties": {
                "dateCount": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "msgCount": {
                    "type": "integer"
                },
                "userCount": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/msg.ActiveUser"
                    }
                }
            }
        },
        "sdkws.GroupInfo": {
            "type": "object",
            "properties": {
                "applyMemberFriend": {
                    "type": "integer"
                },
                "createTime": {
                    "type": "integer"
                },
                "creatorUserID": {
                    "type": "string"
                },
                "ex": {
                    "type": "string"
                },
                "faceURL": {
                    "type": "string"
                },
                "groupID": {
                    "type": "string"
                },
                "groupName": {
                    "type": "string"
                },
                "groupType": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "lookMemberInfo": {
                    "type": "integer"
                },
                "memberCount": {
                    "type": "integer"
                },
                "needVerification": {
                    "type": "integer"
                },
                "notification": {
                    "type": "string"
                },
                "notificationUpdateTime": {
                    "type": "integer"
                },
                "notificationUserID": {
                    "type": "string"
                },
                "ownerUserID": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "sdkws.RequestPagination": {
            "type": "object",
            "properties": {
                "pageNumber": {
                    "type": "integer"
                },
                "showNumber": {
                    "type": "integer"
                }
            }
        },
        "sdkws.UserInfo": {
            "type": "object",
            "properties": {
                "appMangerLevel": {
                    "type": "integer"
                },
                "createTime": {
                    "type": "integer"
                },
                "ex": {
                    "type": "string"
                },
                "faceURL": {
                    "type": "string"
                },
                "globalRecvMsgOpt": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "user.UserRegisterCountReq": {
            "type": "object",
            "properties": {
                "end": {
                    "type": "integer"
                },
                "start": {
                    "type": "integer"
                }
            }
        },
        "user.UserRegisterCountResp": {
            "type": "object",
            "properties": {
                "before": {
                    "type": "integer"
                },
                "count": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "\"Type 'Bearer' followed by a space and JWT token\"",
            "type": "apiKey",
            "name": "token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:10002",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "飞宏IM接口",
	Description:      "飞宏IM接口",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
