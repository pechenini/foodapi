// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/suppliers": {
            "get": {
                "description": "Get suppliers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Suppliers Actions"
                ],
                "summary": "Get suppliers",
                "operationId": "get-suppliers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SupplierCollectionResponse"
                        }
                    }
                }
            }
        },
        "/suppliers/{id}/menu": {
            "get": {
                "description": "Get menu by supplier id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Supplier Actions"
                ],
                "summary": "Get supplier menu by supplier id",
                "operationId": "get-supplier-menu",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Supplier ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MenuResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/suppliers/{id}/menu/{position_id}": {
            "get": {
                "description": "Get menu position by supplier id and position id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Supplier Actions"
                ],
                "summary": "Get supplier menu by supplier id and position id",
                "operationId": "get-supplier-menu-position",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Supplier ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Supplier menu position ID",
                        "name": "position_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Menu"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Menu": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
        },
        "model.WorkingHours": {
            "type": "object",
            "properties": {
                "closing": {
                    "type": "string"
                },
                "opening": {
                    "type": "string"
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "response.MenuResponse": {
            "type": "object",
            "properties": {
                "menu": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Menu"
                    }
                }
            }
        },
        "response.SupplierCollectionResponse": {
            "type": "object",
            "properties": {
                "suppliers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.SupplierResponse"
                    }
                }
            }
        },
        "response.SupplierResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "workingHours": {
                    "$ref": "#/definitions/model.WorkingHours"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
