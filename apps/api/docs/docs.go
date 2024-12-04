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
        "/api/v1/book": {
            "post": {
                "description": "create new book with book info and user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Create new book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID for authentication",
                        "name": "X-User-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Create new book with info and user_id",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validators.CreateBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.CreateBookResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/book/cost": {
            "post": {
                "description": "create new cost with cost information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "costs"
                ],
                "summary": "Create new cost",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID for authentication",
                        "name": "X-User-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Create new cost with cost information",
                        "name": "cost",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validators.CreateCostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.CreateCostResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/book/join": {
            "post": {
                "description": "join the book with book_id and user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Join the book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID for authentication",
                        "name": "X-User-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Join the book with book_id",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validators.JoinBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/book/list": {
            "get": {
                "description": "get book list with user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get book list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID for authentication",
                        "name": "X-User-ID",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.BookModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/book/{book_id}": {
            "get": {
                "description": "get book with book_id (and X-User-ID)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get book by book_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID for authentication",
                        "name": "X-User-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Get book with book_id",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.CombinedBookDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/book/{book_id}/cost/{cost_id}": {
            "get": {
                "description": "get cost detail by book_id and cost_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "costs"
                ],
                "summary": "Get cost detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID for authentication",
                        "name": "X-User-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Cost ID",
                        "name": "cost_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.CostDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/book/{book_id}/leave": {
            "delete": {
                "description": "leave the book with book_id and user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Leave the book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID for authentication",
                        "name": "X-User-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Leave the book with book_id",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "description": "get user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID for authentication",
                        "name": "X-User-ID",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create new user with new device_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "Create new user by device_id",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validators.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.CreateUserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BookMember": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "models.BookModel": {
            "type": "object",
            "properties": {
                "book_description": {
                    "type": "string"
                },
                "book_name": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_modified_at": {
                    "type": "string"
                }
            }
        },
        "models.CostDetail": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "book_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "payers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CostPayerModel"
                    }
                },
                "sharers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CostSharerModel"
                    }
                }
            }
        },
        "models.CostPayerModel": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "cost_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.CostRecordModel": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "book_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "models.CostSharerModel": {
            "type": "object",
            "properties": {
                "cost_id": {
                    "type": "string"
                },
                "share_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_modified_at": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "responses.CombinedBookDetail": {
            "type": "object",
            "properties": {
                "book_description": {
                    "type": "string"
                },
                "book_name": {
                    "type": "string"
                },
                "costs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CostRecordModel"
                    }
                },
                "create_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_modified_at": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BookMember"
                    }
                }
            }
        },
        "responses.CreateBookResponse": {
            "type": "object",
            "properties": {
                "book_description": {
                    "type": "string"
                },
                "book_name": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_modified_at": {
                    "type": "string"
                }
            }
        },
        "responses.CreateCostResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "responses.CreateUserResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "utils.APIResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "validators.CostPayerRequest": {
            "type": "object",
            "required": [
                "amount",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "validators.CostSharerRequest": {
            "type": "object",
            "required": [
                "share_amount",
                "user_id"
            ],
            "properties": {
                "share_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "validators.CreateBookRequest": {
            "type": "object",
            "required": [
                "book_name"
            ],
            "properties": {
                "book_description": {
                    "type": "string"
                },
                "book_name": {
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "validators.CreateCostRequest": {
            "type": "object",
            "required": [
                "amount",
                "book_id",
                "currency",
                "payers",
                "sharers"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "book_id": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "payers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/validators.CostPayerRequest"
                    }
                },
                "sharers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/validators.CostSharerRequest"
                    }
                }
            }
        },
        "validators.CreateUserRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "validators.JoinBookRequest": {
            "type": "object",
            "required": [
                "book_id"
            ],
            "properties": {
                "book_id": {
                    "type": "string"
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
	Title:            "Bill Breaker API",
	Description:      "This is the API of Bill Breaker",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
