package core

import (
	"example.com/t/docs"
	"github.com/gin-gonic/gin"
)

type SwaggerConfig struct {
}

func InitSwagger(r *gin.Engine) {
	docs.SwaggerInfo.Title = "Swagger Gin Template API"
	docs.SwaggerInfo.Description = "Swagger Gin Template API Service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.LeftDelim = "{{"
	docs.SwaggerInfo.RightDelim = "}}"
	docs.SwaggerInfo.InfoInstanceName = "swagger"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.SwaggerTemplate = docTemplate
}

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support"
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
        "/udp/last": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "udp"
                ],
                "summary": "获取最后一条 UDP 消息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.UDPMessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/udp/send": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "udp"
                ],
                "summary": "发送 UDP 消息",
                "parameters": [
                    {
                        "description": "消息体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SendUDPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/ws/broadcast": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "websocket"
                ],
                "summary": "广播 WebSocket 消息",
                "parameters": [
                    {
                        "description": "消息体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.BroadcastRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/ws/count": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "websocket"
                ],
                "summary": "获取当前 WebSocket 连接数",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.WsClientCountResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.BroadcastRequest": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "hello"
                }
            }
        },
        "controller.ErrorResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string",
                    "example": "dial udp: ..."
                },
                "error": {
                    "type": "string",
                    "example": "需要字段 msg"
                }
            }
        },
        "controller.StatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.WsClientCountResponse": {
            "type": "object",
            "properties": {
                "client_count": {
                    "type": "integer",
                    "example": 3
                }
            }
        },
        "main.SendUDPRequest": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "hello udp"
                }
            }
        },
        "main.UDPMessageResponse": {
            "type": "object",
            "properties": {
                "last_msg": {
                    "type": "string",
                    "example": "{...}"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`
