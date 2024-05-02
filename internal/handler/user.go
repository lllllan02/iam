package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lllllan02/iam/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h UserHandler) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, err := h.userService.Login(c, username, password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
