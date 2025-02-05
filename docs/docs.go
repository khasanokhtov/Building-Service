// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/buildings": {
            "get": {
                "description": "Получает список строений с возможностью фильтрации",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Получить список строений",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Город",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Год",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Количество этажей",
                        "name": "floors",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Building"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Добавляет новое строение в базу данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Добавить строение",
                "parameters": [
                    {
                        "description": "Данные строения",
                        "name": "building",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Building"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Building"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Building": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "floors": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Building API",
	Description:      "API для управления строениями",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
