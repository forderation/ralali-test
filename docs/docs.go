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
        "/cakes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cakes"
                ],
                "summary": "GetCakes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "default page is at page 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "maximum value is 100",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetCakesResponse"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cakes"
                ],
                "summary": "CreateCake",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ApiMutationCakePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CakeMutationResponse"
                        }
                    }
                }
            }
        },
        "/cakes/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cakes"
                ],
                "summary": "GetCake",
                "parameters": [
                    {
                        "type": "string",
                        "description": "param id (cake record)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CakeResponse"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cakes"
                ],
                "summary": "UpdateCake",
                "parameters": [
                    {
                        "type": "string",
                        "description": "param id (cake record)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ApiMutationCakePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CakeMutationResponse"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cakes"
                ],
                "summary": "DeleteCake",
                "parameters": [
                    {
                        "type": "string",
                        "description": "param id (cake record)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CakeDeleteResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ApiMutationCakePayload": {
            "type": "object",
            "required": [
                "rating",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.CakeDeleteResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.CakeMutationResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.CakeResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.GetCakesResponse": {
            "type": "object",
            "properties": {
                "cakes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CakeResponse"
                    }
                },
                "meta": {
                    "$ref": "#/definitions/model.MetaPagination"
                }
            }
        },
        "model.MetaPagination": {
            "type": "object",
            "properties": {
                "page_count": {
                    "type": "integer"
                },
                "total_data": {
                    "type": "integer"
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
