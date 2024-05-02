//go:build wireinject

package wire

import (
	"net/http"

	"github.com/google/wire"
	"github.com/lllllan02/iam/internal/config"
	"github.com/lllllan02/iam/internal/handler"
	"github.com/lllllan02/iam/internal/repository"
	"github.com/lllllan02/iam/internal/server"
	"github.com/lllllan02/iam/internal/service"
)

var configSet = wire.NewSet(config.NewConfig)

var serverSet = wire.NewSet(server.NewServer)

var handlerSet = wire.NewSet(handler.NewUserHandler)

var serviceSet = wire.NewSet(service.NewUserService)

var repoSet = wire.NewSet(
	repository.NewRepo,
	repository.NewTransaction,
	repository.NewUserRepo,
)

func Init() *http.Server {
	wire.Build(
		configSet,
		serverSet,
		handlerSet,
		serviceSet,
		repoSet,
	)

	return new(http.Server)
}
