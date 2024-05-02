package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/iam/internal/handler"
)

func NewServer(
	userHandler *handler.UserHandler,
) *http.Server {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "Hello, World!") })
	engine.GET("/login", userHandler.Login)

	return &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
}
