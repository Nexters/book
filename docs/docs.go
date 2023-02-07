// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/books": {
            "get": {
                "description": "사용자가 등록한 모든 책을 조회하는 API. TODO: 읽을책/완독 구분해 가져오게 할 예정",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "사용자가 등록한 모든 책을 조회하는 API",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "default = true",
                        "name": "isReading",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payloads.FindAllBooksPayload"
                        }
                    }
                }
            },
            "post": {
                "description": "책의 ISBN, 제목, userId를 body로 제공하면 읽을책으로 등록하는 API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "읽을 책을 등록하는 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "CreateBookParam{}",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateBookParam"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Book"
                        }
                    }
                }
            }
        },
        "/books/search": {
            "get": {
                "description": "Naver API를 이용해 책을 검색하게 하는 API query string으로 title을 넘기면 검색 결과를 반환.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "책 검색 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "미움받을 용기",
                        "name": "title",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/search.SearchItem"
                            }
                        }
                    }
                }
            }
        },
        "/books/{bookId}": {
            "get": {
                "description": "bookID로 유저의 책과 그에 대한 모든 메모를 조회하는 API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "bookID 혹은 ISBN으로 책과 모든 메모 조회 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "12345678",
                        "name": "bookId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "true",
                        "name": "isbn",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "comment",
                        "name": "category",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Book"
                        }
                    }
                }
            }
        },
        "/memos": {
            "post": {
                "description": "특정 사용자가 특정 책에 대해 메모를 작성하는 API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memo"
                ],
                "summary": "메모 추가 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "service.CreateMemoParam{}",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.CreateMemoParam"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Memo"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Authorization header의 bearer token을 이용해 사용자 통계를 조회함",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "사용자 통계 조회 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payloads.UserStatPayload"
                        }
                    }
                }
            }
        },
        "/users/token": {
            "get": {
                "description": "API를 호출하면 UUID를 token으로 발급함. local storage에 저장해두고 userId로 사용하면 됨.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "사용자 추가 API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CreateBookParam": {
            "type": "object",
            "required": [
                "ISBN",
                "title"
            ],
            "properties": {
                "ISBN": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "entity.Book": {
            "type": "object",
            "properties": {
                "ISBN": {
                    "type": "string"
                },
                "author": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
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
                "isReading": {
                    "type": "boolean"
                },
                "memos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Memo"
                    }
                },
                "price": {
                    "type": "string"
                },
                "publisher": {
                    "type": "string"
                },
                "shopLink": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "entity.Memo": {
            "type": "object",
            "properties": {
                "bookID": {
                    "type": "integer"
                },
                "category": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Book"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "payloads.FindAllBooksPayload": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/payloads.FindBookPayload"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "payloads.FindBookPayload": {
            "type": "object",
            "properties": {
                "ISBN": {
                    "type": "string"
                },
                "author": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
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
                "isReading": {
                    "type": "boolean"
                },
                "memoCount": {
                    "type": "integer"
                },
                "memos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Memo"
                    }
                },
                "price": {
                    "type": "string"
                },
                "publisher": {
                    "type": "string"
                },
                "shopLink": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "payloads.UserStatPayload": {
            "type": "object",
            "properties": {
                "duration": {
                    "type": "integer"
                },
                "memoCount": {
                    "type": "integer"
                },
                "readCount": {
                    "type": "integer"
                }
            }
        },
        "search.SearchItem": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "discount": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "isbn": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "pubdate": {
                    "type": "string"
                },
                "publisher": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "service.CreateMemoParam": {
            "type": "object",
            "required": [
                "bookId",
                "category",
                "text"
            ],
            "properties": {
                "bookId": {
                    "type": "integer"
                },
                "category": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
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
