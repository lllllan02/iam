//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lllllan02/iam/pkg/app"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/log"
)

func newApp() *app.App {
	return app.NewApp(
		app.WithName("iam-server"),
	)
}

func NewWire(*config.Config, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		newApp,
	))
}
