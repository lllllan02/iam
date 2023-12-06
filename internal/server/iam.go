package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/iam/api"
	"github.com/lllllan02/iam/internal/handler"
	"github.com/lllllan02/iam/internal/middleware"
	"github.com/lllllan02/iam/pkg/code"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/errors"
	"github.com/lllllan02/iam/pkg/log"
	"github.com/lllllan02/iam/pkg/resp"
	"github.com/lllllan02/iam/pkg/server/http"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewIAMServer(
	conf *config.Config,
	logger *log.Logger,
	userHandler *handler.UserHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.Server.Host),
		http.WithServerPort(conf.Server.HttpPort),
	)

	// use ginSwagger middleware to serve the API docs
	api.SwaggerInfo.Title = "IAM 身份识别与管理系统"
	api.SwaggerInfo.Version = "0.1"
	api.SwaggerInfo.Description = "基于 Go 开发的身份识别与管理系统，可用于第三方登录及资源/权限管理"
	api.SwaggerInfo.Host = fmt.Sprint("localhost:", conf.Server.HttpPort)
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// middleware
	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
	)

	// router
	s.GET("/ping", func(ctx *gin.Context) { resp.JsonResponse(ctx, nil, ":) pong") })
	s.GET("/error", func(ctx *gin.Context) {
		resp.JsonResponse(ctx, errors.WithCode(code.C_ExampleProject_ExampleModule_ExampleErr, "error"), nil)
	})

	s.POST("/register", userHandler.Register)

	return s
}
