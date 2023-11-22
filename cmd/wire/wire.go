//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lllllan02/iam/internal/data"
	"github.com/lllllan02/iam/internal/handler"
	"github.com/lllllan02/iam/internal/server"
	"github.com/lllllan02/iam/internal/service"
	"github.com/lllllan02/iam/pkg/app"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/log"
	"github.com/lllllan02/iam/pkg/server/http"
)

var serverSet = wire.NewSet(
	server.NewIAMServer,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var dataSet = wire.NewSet(
	data.NewDB,
	data.NewRedis,
	data.NewData,
	data.NewTransaction,
	data.NewUserData,
)

func newApp(iamServer *http.Server) *app.App {
	return app.NewApp(
		app.WithServer(iamServer),
		app.WithName("iam-server"),
	)
}

func NewWire(*config.Config, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		newApp,
		serverSet,
		handlerSet,
		serviceSet,
		dataSet,
	))
}
