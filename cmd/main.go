package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/lllllan02/iam/cmd/wire"
	"github.com/lllllan02/iam/pkg/config"
)

func main() {
	var confPath = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()

	conf := config.NewConfig(*confPath)
	fmt.Printf("conf: %v\n", conf)

	// TODO log

	app, cleanup, err := wire.NewWire(conf)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
