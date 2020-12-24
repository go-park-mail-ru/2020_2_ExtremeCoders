// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "martin7.heinz@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/letter": {
            "post": {
                "description": "send {to:'receiver', theme:'theme', text:'letter content'}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "send letter",
                "operationId": "send-letter",
                "parameters": [
                    {
                        "description": "Letter ID",
                        "name": "letter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LetterModel.Letter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "delete letter {id:10}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete letter",
                "operationId": "delete-letter",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Letter ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/letter/by/{what}/{value}": {
            "get": {
                "description": "поиск по всем файлам\nget letter/by/{what}/{value} - what может быть равен\n(id, sender, receiver, theme, text, date_time, directory_recv, directory_send)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Search in letter",
                "operationId": "all-search",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search type",
                        "name": "what",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "search template",
                        "name": "value",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/LetterModel.Letter"
                            }
                        }
                    }
                }
            }
        },
        "/letter/{similar}": {
            "get": {
                "description": "get letter/{similar} - поиск по всем письмам",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Search in letter",
                "operationId": "search-search",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search template",
                        "name": "similar",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/session": {
            "post": {
                "description": "get user by ID body {email:\"email@mail.ru\", password: \"password\"}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "login",
                "operationId": "get-session",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UserModel.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "logout: remove session cookies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "logout",
                "operationId": "get-session",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "get user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "changeUserProfile",
                "operationId": "post-user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UserModel.User"
                        }
                    }
                }
            },
            "post": {
                "description": "user registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "signup",
                "operationId": "post-user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UserModel.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UserModel.User"
                        }
                    }
                }
            }
        },
        "/user/avatar": {
            "get": {
                "description": "get user avatar",
                "consumes": [
                    "application/json"
                ],
                "summary": "get user avatar",
                "operationId": "get-user-avatar",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/user/folders/recived/folderName": {
            "put": {
                "description": "Переименовать папку  user/folders/{recived/sended}/folderName {oldName:\"oldName\", newName:\"newName\"}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Rename Folder",
                "operationId": "RenameFolder",
                "parameters": [
                    {
                        "description": "folder name",
                        "name": "folderName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FolderDelivery.Folder"
                        }
                    },
                    {
                        "description": "Letter id",
                        "name": "letterId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "description": "добавить папку в полученные (отправленные) post user/folders/{recived/sended}/folderName {folderName:\"folderName\"}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add folder",
                "operationId": "AddFolder",
                "parameters": [
                    {
                        "description": "folder name",
                        "name": "folderName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FolderDelivery.Folder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "удалить папку delete user/folders/{recived/sended}/folderName {folderName:\"folderName\"}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove Folder",
                "operationId": "RemoveFolder",
                "parameters": [
                    {
                        "description": "folder name",
                        "name": "folderName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FolderDelivery.Folder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/folders/recived/folderName/letter": {
            "put": {
                "description": "добавить писмо в папку post user/folders/{recived/sended}/folderName {folderName:\"folderName\", letterID: id}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add letter in folder",
                "operationId": "AddLetterInFolder",
                "parameters": [
                    {
                        "description": "folder name",
                        "name": "folderName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FolderDelivery.Folder"
                        }
                    },
                    {
                        "description": "Letter id",
                        "name": "letterId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/folders/sended/folderName/letter": {
            "delete": {
                "description": "Удалить письмо из папки user/folders/{recived/sended}/folderName\n/user/folders/sended/folderName/letter body{letterID:Id}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove Letter from Folder",
                "operationId": "RemoveLetterInFolder",
                "parameters": [
                    {
                        "description": "folder name",
                        "name": "folderName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/FolderDelivery.Folder"
                        }
                    },
                    {
                        "description": "Letter id",
                        "name": "letterId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/folders/{recived/sended}": {
            "get": {
                "description": "user/folders/{recived/sended} - список папок в отправленных (полученных) письмах",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "getFolderList",
                "operationId": "GetFolderList",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/FolderDelivery.FolderList"
                        }
                    }
                }
            }
        },
        "/user/folders/{recived/sended}/{folderName}/{limit}/{offset}": {
            "get": {
                "description": "письма из папки в полученых (отправленных) user/foders/{recived/sended}/{folderName}/{limit}/{offset}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get letters by folder",
                "operationId": "GetLettersByFolder",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/FolderDelivery.LetterList"
                        }
                    }
                }
            }
        },
        "/user/letter/received/{limit}/{offset}": {
            "get": {
                "description": "get user/letter/sent/{limit}/{offset} - получить полученные письма",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get received letter",
                "operationId": "get-received-letter",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/LetterModel.Letter"
                            }
                        }
                    }
                }
            }
        },
        "/user/letter/sent/{limit}/{offset}": {
            "get": {
                "description": "get user/letter/sent/{limit}/{offset} - получить отправленные письма",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get sended letter",
                "operationId": "get-send-letter",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/LetterModel.Letter"
                            }
                        }
                    }
                }
            }
        },
        "/watch/letter": {
            "put": {
                "description": "отметить письмо как прочитанное/непрочитанное /watch/letter {id:'id'}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "set watch togle",
                "operationId": "watch-letter",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "letter id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "FolderDelivery.Folder": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "FolderDelivery.FolderList": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "folders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/FolderDelivery.Folder"
                    }
                }
            }
        },
        "FolderDelivery.LetterList": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "letter": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/LetterModel.Letter"
                    }
                }
            }
        },
        "LetterModel.Letter": {
            "type": "object",
            "properties": {
                "dateTime": {
                    "type": "integer"
                },
                "directoryRecv": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "isWatched": {
                    "type": "boolean"
                },
                "receiver": {
                    "type": "string"
                },
                "sender": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "theme": {
                    "type": "string"
                }
            }
        },
        "UserModel.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "img": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "errors.AnswerGet": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/UserModel.User"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Mailer Swagger API",
	Description: "Swagger API for Golang Project Blueprint.",
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