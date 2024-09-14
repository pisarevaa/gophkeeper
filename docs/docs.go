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
        "/api/data": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "Get all data",
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.DataResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/api/data/binary": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "Add text data",
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.DataResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity (body)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/api/data/binary/{dataID}": {
            "put": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "Update binary data",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Data ID",
                        "name": "dataId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.DataResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity (query or body)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/api/data/text": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "Add text data",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddTextData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.DataResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity (body)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/api/data/text/{dataID}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "Update text data",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Data ID",
                        "name": "dataId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddTextData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.DataResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity (query or body)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/api/data/{dataID}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "Get data by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Data ID",
                        "name": "dataId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.DataResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "Data is not found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity (query)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "Delete data",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Data ID",
                        "name": "dataId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.DataResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity (query)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.TokenResponse"
                        }
                    },
                    "401": {
                        "description": "Incorrect password",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "Email is not found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Regiser user",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    },
                    "409": {
                        "description": "Email is already used",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AddTextData": {
            "type": "object",
            "required": [
                "data",
                "name"
            ],
            "properties": {
                "data": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 250
                }
            }
        },
        "model.DataResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/model.DataTypeEnum"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.DataTypeEnum": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-comments": {
                "BinaryType": "binary",
                "TextType": "text",
                "TypeUnknown": "unknown"
            },
            "x-enum-varnames": [
                "TypeUnknown",
                "TextType",
                "BinaryType"
            ]
        },
        "model.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.RegisterUser": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 250
                },
                "password": {
                    "type": "string",
                    "maxLength": 250
                }
            }
        },
        "model.TokenResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "model.UserResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Gophkeeper API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
