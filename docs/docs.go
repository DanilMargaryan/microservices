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
            "name": "API Support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/beverage": {
            "post": {
                "description": "Добавляет новый напиток в базу данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "beverages"
                ],
                "summary": "Добавить новый напиток",
                "parameters": [
                    {
                        "description": "Данные напитка",
                        "name": "beverage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/storage.Beverage"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Напиток успешно добавлен!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка при парсинге тела запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при добавлении напитка в базу данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/beverage/{id}": {
            "get": {
                "description": "Возвращает данные напитка по его идентификатору",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "beverages"
                ],
                "summary": "Получить напиток по id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID напитка",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/storage.Beverage"
                        }
                    },
                    "404": {
                        "description": "Напиток не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при получении данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/beverages": {
            "get": {
                "description": "Возвращает список всех напитков из базы данных",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "beverages"
                ],
                "summary": "Получить все напитки",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/storage.Beverage"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка при получении данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "storage.Beverage": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Beverage API",
	Description:      "API для управления напитками.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
