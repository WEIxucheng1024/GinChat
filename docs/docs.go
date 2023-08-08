// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/test2": {
            "get": {
                "tags": [
                    "首页"
                ],
                "responses": {
                    "200": {
                        "description": "welcome",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/createUser": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "新增用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "userName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "passWord",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "phone",
                        "name": "phone",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": \"200\", \"message\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "userName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": \"200\", \"message\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserByPhone": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "根据手机号获取用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": \"200\", \"message\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserByUEmail": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "根据邮箱获取用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": \"200\", \"message\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserByUserName": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "根据userName获取用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "userName",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": \"200\", \"message\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "所有用户",
                "responses": {
                    "200": {
                        "description": "code\": \"200\", \"message\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/loginUser": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "登录用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "userName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "passWord",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": \"200\", \"message\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/updateUser": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "userName",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "名称",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "passWord",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": \"200\", \"message\": \"Success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
