package app

import (
	"context"
	"expvar"
	"test/internal/config"
)

func Run() {
	ctx := context.Background()

	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	expvar.Publish("config", conf)

	app, err := NewApp(ctx, conf)
	if err != nil {
		panic(err)
	}

	app.Run(ctx)
}
