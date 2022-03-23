package controllers

import (
	postcontrollers "github.com/maoaeri/openapi/pkg/controllers/post"
	usercontrollers "github.com/maoaeri/openapi/pkg/controllers/user"
)

type Controllers struct {
	UserController *usercontrollers.UserController
	PostController *postcontrollers.PostController
}
