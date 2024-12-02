// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Current API version",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Version"
                ],
                "summary": "Returns the API version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.Message"
                        }
                    }
                }
            }
        },
        "/v1/conversion": {
            "get": {
                "description": "Convert a specified amount from one currency to another using the latest exchange rates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "conversion"
                ],
                "summary": "Convert a amount",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Currency code to convert from",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Currency code to convert to",
                        "name": "to",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Amount to be converted",
                        "name": "amount",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/conversion.ConversionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.Message"
                        }
                    }
                }
            }
        },
        "/v1/currency": {
            "get": {
                "description": "Show all currencies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "currency"
                ],
                "summary": "List currencies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/currency.Currency"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.Message"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Currency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "currency"
                ],
                "summary": "Create Currency",
                "parameters": [
                    {
                        "description": "Currency Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/currency.CreateData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/currency.Currency"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.Message"
                        }
                    }
                }
            }
        },
        "/v1/currency/{id}": {
            "put": {
                "description": "Update a Currency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "currency"
                ],
                "summary": "Update Currency",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the Currency",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Currency Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/currency.UpdateData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/currency.Currency"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.Message"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Currency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "currency"
                ],
                "summary": "Delete Currency",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the Currency",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.Message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "conversion.ConversionResponse": {
            "description": "Response from conversion",
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "currency.CreateData": {
            "description": "Request to Create a new Currency",
            "type": "object",
            "required": [
                "code",
                "currency_rate",
                "name"
            ],
            "properties": {
                "backing_currency": {
                    "type": "boolean",
                    "default": false
                },
                "code": {
                    "type": "string"
                },
                "currency_rate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "currency.Currency": {
            "type": "object",
            "properties": {
                "backing_currency": {
                    "type": "boolean"
                },
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "currency_rate": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "currency.UpdateData": {
            "description": "Request to Update a Currency",
            "type": "object",
            "required": [
                "code",
                "currency_rate",
                "name"
            ],
            "properties": {
                "backing_currency": {
                    "type": "boolean",
                    "default": false
                },
                "code": {
                    "type": "string"
                },
                "currency_rate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "gin.Message": {
            "description": "Standard response to errors",
            "type": "object",
            "properties": {
                "message": {
                    "description": "This is a json message",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "status": "message"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8085",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Conversor de Moedas",
	Description:      "API para conversão de moedas",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
