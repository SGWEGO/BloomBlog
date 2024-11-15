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
        "/bloomblog/relation/action": {
            "post": {
                "description": "Handles actions like follow, unfollow, or other user relation actions.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Relation"
                ],
                "summary": "Perform a relation action",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User authentication token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The ID of the user to perform the action on",
                        "name": "to_user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "1",
                            "2"
                        ],
                        "type": "string",
                        "example": "1",
                        "description": "The type of action to perform",
                        "name": "action_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Action completed successfully",
                        "schema": {
                            "$ref": "#/definitions/relation.BloomblogRelationActionResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input parameters",
                        "schema": {
                            "$ref": "#/definitions/relation.BloomblogRelationActionResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/relation.BloomblogRelationActionResponse"
                        }
                    }
                }
            }
        },
        "/bloomblog/relation/followerlist": {
            "get": {
                "description": "Retrieves the list of users that follow the specified user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Relation"
                ],
                "summary": "Get follower list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The ID of the user to retrieve the follower list for",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User authentication token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved follower list",
                        "schema": {
                            "$ref": "#/definitions/relation.BloomblogRelationFollowerListResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input parameters",
                        "schema": {
                            "$ref": "#/definitions/relation.BloomblogRelationFollowerListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/relation.BloomblogRelationFollowerListResponse"
                        }
                    }
                }
            }
        },
        "/bloomblog/relation/followlist": {
            "get": {
                "description": "Retrieves the list of users that a specific user is following.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Relation"
                ],
                "summary": "Get following list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The ID of the user to retrieve the following list for",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User authentication token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved following list",
                        "schema": {
                            "$ref": "#/definitions/relation.BloomblogRelationFollowListResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input parameters",
                        "schema": {
                            "$ref": "#/definitions/relation.BloomblogRelationFollowListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/relation.BloomblogRelationFollowListResponse"
                        }
                    }
                }
            }
        },
        "/bloomblog/user/getuserbyid": {
            "post": {
                "description": "Get user information by ID and token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User by ID",
                "parameters": [
                    {
                        "description": "User ID and token",
                        "name": "userVar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.BloomBlogUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errno.ErrNo"
                        }
                    }
                }
            }
        },
        "/bloomblog/user/login": {
            "post": {
                "description": "Authenticate user with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "User login data",
                        "name": "loginParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserRegisterParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.BloomBlogUserRegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errno.ErrNo"
                        }
                    }
                }
            }
        },
        "/bloomblog/user/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Registration",
                "parameters": [
                    {
                        "description": "User registration data",
                        "name": "registerParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserRegisterParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.BloomBlogUserRegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errno.ErrNo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errno.ErrNo": {
            "type": "object",
            "properties": {
                "errCode": {
                    "type": "integer"
                },
                "errMsg": {
                    "type": "string"
                }
            }
        },
        "handlers.UserParam": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "用户鉴权token",
                    "type": "string"
                },
                "user_id": {
                    "description": "用户id",
                    "type": "integer"
                }
            }
        },
        "handlers.UserRegisterParam": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "用户密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "relation.BloomblogRelationActionResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "description": "0-success，others-failure",
                    "type": "integer"
                },
                "status_msg": {
                    "description": "return statement description",
                    "type": "string"
                }
            }
        },
        "relation.BloomblogRelationFollowListResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "description": "0-success others-failure",
                    "type": "integer"
                },
                "status_msg": {
                    "description": "status description",
                    "type": "string"
                },
                "user_list": {
                    "description": "user list",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.User"
                    }
                }
            }
        },
        "relation.BloomblogRelationFollowerListResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "user_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.User"
                    }
                }
            }
        },
        "user.BloomBlogUserRegisterResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "description": "status code, 0 for success, others for failure",
                    "type": "integer"
                },
                "status_msg": {
                    "description": "description for status",
                    "type": "string"
                },
                "token": {
                    "description": "credential token",
                    "type": "string"
                },
                "user_id": {
                    "description": "user id",
                    "type": "integer"
                }
            }
        },
        "user.BloomBlogUserResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "description": "status code, 0 for success, others for failure",
                    "type": "integer"
                },
                "status_msg": {
                    "description": "description for status",
                    "type": "string"
                },
                "user": {
                    "description": "user information",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.User"
                        }
                    ]
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "follow_count": {
                    "description": "follow count",
                    "type": "integer"
                },
                "follower_count": {
                    "description": "followers count",
                    "type": "integer"
                },
                "id": {
                    "description": "user id",
                    "type": "integer"
                },
                "is_follow": {
                    "description": "true-followed，false-not followed",
                    "type": "boolean"
                },
                "name": {
                    "description": "user name",
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
