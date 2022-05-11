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
        "/queryEventByBlockRange": {
            "post": {
                "description": "queries event by block range",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "flow-event-fetcher"
                ],
                "summary": "queries event by block range",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.QueryEventByBlockRangeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v1.QueryEventByBlockRangeResponseEvent"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ResponseError"
                        }
                    }
                }
            }
        },
        "/queryLatestBlockHeight": {
            "get": {
                "description": "queries the latest block height",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "flow-event-fetcher"
                ],
                "summary": "queries the latest block height",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.QueryLatestBlockHeightResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ResponseError"
                        }
                    }
                }
            }
        },
        "/syncSpork": {
            "get": {
                "description": "sync spork",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "flow-event-fetcher"
                ],
                "summary": "sync spork",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.QueryLatestBlockHeightResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ResponseError"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "get version",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "flow-event-fetcher"
                ],
                "summary": "get version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.VersionResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "timestamppb.Timestamp": {
            "type": "object",
            "properties": {
                "nanos": {
                    "description": "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.",
                    "type": "integer"
                },
                "seconds": {
                    "description": "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.",
                    "type": "integer"
                }
            }
        },
        "v1.QueryEventByBlockRangeRequest": {
            "type": "object",
            "properties": {
                "end": {
                    "type": "integer"
                },
                "event": {
                    "type": "string"
                },
                "start": {
                    "type": "integer"
                }
            }
        },
        "v1.QueryEventByBlockRangeResponseEvent": {
            "type": "object",
            "properties": {
                "blockId": {
                    "type": "integer"
                },
                "eventID": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "timestamp": {
                    "$ref": "#/definitions/timestamppb.Timestamp"
                },
                "transactionId": {
                    "type": "string"
                },
                "transactionIndex": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "values": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.QueryEventByBlockRangeResponseValue"
                    }
                }
            }
        },
        "v1.QueryEventByBlockRangeResponseValue": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "v1.QueryLatestBlockHeightResponse": {
            "type": "object",
            "properties": {
                "latestBlockHeight": {
                    "type": "integer"
                }
            }
        },
        "v1.VersionResponse": {
            "type": "object",
            "properties": {
                "backendMode": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
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
	Version:     "1.0.1",
	Host:        "localhost:8989",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "flow-event-fetcher API",
	Description: "flow-event-fetcher interface documentation",
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
	swag.Register("swagger", &s{})
}
