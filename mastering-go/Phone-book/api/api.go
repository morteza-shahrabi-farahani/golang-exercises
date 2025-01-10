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
        "/": {
            "get": {
                "description": "Respond with a 404 error for unknown routes",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "misc"
                ],
                "summary": "Handle invalid routes",
                "responses": {
                    "404": {
                        "description": "This page does not exist",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/delete/{id}": {
            "delete": {
                "description": "Delete an entry by its ID",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "phonebook"
                ],
                "summary": "Delete a phonebook entry",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Entry ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/insert": {
            "post": {
                "description": "Add a new entry to the phonebook",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "phonebook"
                ],
                "summary": "Insert a new phonebook entry",
                "parameters": [
                    {
                        "description": "Phonebook Entry",
                        "name": "entry",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/phonebook.Entry"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/phonebook.InsertResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "Get all phonebook entries",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "phonebook"
                ],
                "summary": "List phonebook entries",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/phonebook.ListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/search": {
            "get": {
                "description": "Search for an entry by phone number",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "phonebook"
                ],
                "summary": "Search phonebook entries",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Phone number to search",
                        "name": "phone-number",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/phonebook.Entry"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "phonebook.Entry": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "phonebook.InsertResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "phonebook.ListResponse": {
            "type": "object",
            "properties": {
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/phonebook.Entry"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}