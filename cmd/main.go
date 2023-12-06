package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/lllllan02/iam/cmd/wire"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/log"
	"go.uber.org/zap"
)

func main() {
	var confPath = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()

	conf := config.NewConfig(*confPath)
	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	logger.Info("server start", zap.String("host", fmt.Sprintf("http://%s:%d", conf.Server.Host, conf.Server.HttpPort)))
	logger.Info("api docs", zap.String("addr", fmt.Sprintf("http://%s:%d/swagger/index.html", conf.Server.Host, conf.Server.HttpPort)))

	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
