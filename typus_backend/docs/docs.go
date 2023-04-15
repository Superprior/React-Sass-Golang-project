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
        "/auth_admin/": {
            "post": {
                "description": "Check for admin rights and give the permission",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authorize admin user",
                "parameters": [
                    {
                        "description": "Provided password",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apiserver.PasswordBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Autorized successfully",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Password was not provided",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    },
                    "401": {
                        "description": "Wrong password provided",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Could not handle the request (server error)",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    }
                }
            }
        },
        "/languages": {
            "get": {
                "description": "Get the complete Languages list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Language"
                ],
                "summary": "Get all Languages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Language"
                            }
                        }
                    },
                    "500": {
                        "description": "Could not query the request or encode JSON",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    }
                }
            }
        },
        "/samples": {
            "get": {
                "description": "Get the complete Samples list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sample"
                ],
                "summary": "Get all Samples",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Sample"
                            }
                        }
                    },
                    "500": {
                        "description": "Could not query the request or encode JSON",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Sample instance",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sample"
                ],
                "summary": "Create Sample",
                "parameters": [
                    {
                        "description": "Provided data for creating Sample",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apiserver.PostSampleBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Returns id of the created Sample",
                        "schema": {
                            "$ref": "#/definitions/apiserver.IdResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid data provided",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Could not create Sample instance",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    }
                }
            }
        },
        "/samples/{id}": {
            "get": {
                "description": "Retvieve a sample instance by provided ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sample"
                ],
                "summary": "Get Sample by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sample ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Sample"
                        }
                    },
                    "400": {
                        "description": "Invalid ID provided or no sample with such ID",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Could not encode JSON",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Sample instance. Available only for admin user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sample"
                ],
                "summary": "Delete Sample",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sample ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns id of the deleted Sample",
                        "schema": {
                            "$ref": "#/definitions/apiserver.IdResponse"
                        }
                    },
                    "400": {
                        "description": "invalid id provided",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Could not delete Sample instance",
                        "schema": {
                            "$ref": "#/definitions/apiserver.MessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apiserver.IdResponse": {
            "description": "Object that is returned in endpoints where ID needs to be returned",
            "type": "object"
        },
        "apiserver.MessageResponse": {
            "description": "Object that is returned when the actual body is empty",
            "type": "object"
        },
        "apiserver.PasswordBody": {
            "description": "Object that needs to be provided for authorizing a user",
            "type": "object",
            "properties": {
                "pwd": {
                    "type": "string"
                }
            }
        },
        "apiserver.PostSampleBody": {
            "description": "Object that stores info about Sample that need to be created",
            "type": "object",
            "properties": {
                "Content": {
                    "type": "string"
                },
                "LangSlug": {
                    "type": "string"
                },
                "Title": {
                    "type": "string"
                }
            }
        },
        "model.Language": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "slug": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Sample": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "langSlug": {
                    "type": "string"
                },
                "title": {
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
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Backend server of Typus",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
