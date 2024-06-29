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
        "/checkin": {
            "post": {
                "description": "Create a check-in using email, event and user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "checkin"
                ],
                "summary": "Create a check-in",
                "parameters": [
                    {
                        "description": "Create CheckIn Request",
                        "name": "create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCheckInRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCheckInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            }
        },
        "/checkin/email/{email}": {
            "get": {
                "description": "Find check-ins by email",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "checkin"
                ],
                "summary": "Find check-ins by email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FindByEmailCheckInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            }
        },
        "/checkin/{userId}": {
            "get": {
                "description": "Find check-ins by user_id",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "checkin"
                ],
                "summary": "Find check-ins by user_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FindByUserIdCheckInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apperror.AppError": {
            "type": "object",
            "properties": {
                "httpCode": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "dto.CheckIn": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "event": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.CreateCheckInRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "event": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.CreateCheckInResponse": {
            "type": "object",
            "properties": {
                "checkin": {
                    "$ref": "#/definitions/dto.CheckIn"
                }
            }
        },
        "dto.FindByEmailCheckInResponse": {
            "type": "object",
            "properties": {
                "checkins": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CheckIn"
                    }
                }
            }
        },
        "dto.FindByUserIdCheckInResponse": {
            "type": "object",
            "properties": {
                "checkins": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CheckIn"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3001",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "RPKM67 API",
	Description:      "the RPKM67 API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}