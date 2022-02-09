package main

import (
	"flag"

	"github.com/A-SoulFan/acao-homework/internal/app/admin-api/server/gin"
)

var configFile = flag.String("f", "config/admin-api.yml", "set config file which viper will loading.")

func main() {
	flag.Parse()

	server, err := gin.InitServer(*configFile)
	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}

	server.AwaitSignal()
}
