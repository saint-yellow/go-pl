// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/api/user/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/serialization.Result"
                        }
                    },
                    "400": {
                        "description": "{\"message\":\"err\",\"code\":1}",
                        "schema": {
                            "$ref": "#/definitions/serialization.Result"
                        }
                    }
                }
            }
        },
        "/api/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户名, 密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.RegisterUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/serialization.Result"
                        }
                    },
                    "400": {
                        "description": "{\"message\":\"err\",\"code\":1}",
                        "schema": {
                            "$ref": "#/definitions/serialization.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serialization.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "service.LoginUser": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 5,
                    "example": "saint-yellow"
                },
                "userName": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3,
                    "example": "saint-yellow"
                }
            }
        },
        "service.RegisterUser": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 5,
                    "example": "saint-yellow"
                },
                "userName": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3,
                    "example": "saint-yellow"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:3000",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "RESTful API collection for Memorandum web app",
	Description: "This is a sample server celler server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
