package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lllllan02/iam/internal/handler"
	"github.com/lllllan02/iam/internal/middleware"
	"github.com/lllllan02/iam/pkg/code"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/errors"
	"github.com/lllllan02/iam/pkg/log"
	"github.com/lllllan02/iam/pkg/resp"
	"github.com/lllllan02/iam/pkg/server/http"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
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

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		// ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

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
