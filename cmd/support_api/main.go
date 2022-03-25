package main

import (
	"flag"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api"
	"github.com/A-SoulFan/acao-homework/internal/launch"
	"github.com/A-SoulFan/acao-homework/internal/pkg/config"
)

var configFile = flag.String("f", "config/support-api.yml", "set config file which viper will loading.")

func main() {
	flag.Parse()

	launch.Launch()

	app, err := support_api.InitApp(config.Path(*configFile))
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
