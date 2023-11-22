package handler

import "github.com/lllllan02/iam/internal/service"

type UserHandler struct {
	*Handler

	userService *service.UserService
}

func NewUserHandler(handler *Handler, userService *service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}
