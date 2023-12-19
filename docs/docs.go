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
        "/solver": {
            "post": {
                "description": "provide letter sets and get suitable variants",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "solver"
                ],
                "summary": "5 letters quiz solver",
                "parameters": [
                    {
                        "description": "Existing letters",
                        "name": "SolverRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SolverRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of suitable words",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.SolverRequest": {
            "type": "object",
            "properties": {
                "absent": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "в",
                        "г"
                    ]
                },
                "ordered": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "3": "о",
                        "5": "е"
                    }
                },
                "unordered": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "а",
                        "б"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
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
