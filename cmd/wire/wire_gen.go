// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lllllan02/iam/internal/config"
	"github.com/lllllan02/iam/internal/handler"
	"github.com/lllllan02/iam/internal/repository"
	"github.com/lllllan02/iam/internal/server"
	"github.com/lllllan02/iam/internal/service"
	"net/http"
)

// Injectors from wire.go:

func Init() *http.Server {
	configConfig := config.NewConfig()
	repo := repository.NewRepo(configConfig)
	userRepo := repository.NewUserRepo(repo)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	httpServer := server.NewServer(userHandler)
	return httpServer
}

// wire.go:

var configSet = wire.NewSet(config.NewConfig)

var serverSet = wire.NewSet(server.NewServer)

var handlerSet = wire.NewSet(handler.NewUserHandler)

var serviceSet = wire.NewSet(service.NewUserService)

var repoSet = wire.NewSet(repository.NewRepo, repository.NewTransaction, repository.NewUserRepo)