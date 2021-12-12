package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/A-SoulFan/acao-homework/internal/app/support-api/config"
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/server/gin/router"
	"github.com/gin-gonic/gin"
)

var configFile = flag.String("f", "config/open.yml", "set config file which viper will loading.")

func main() {
	flag.Parse()
	c := loadConfig()

	svc := svcCtx.NewServiceContext(c)
	defer svc.Stop()

	r := gin.Default()
	router.InitRouter(r, svc)

	_ = r.Run(c.App.Port)
}

func loadConfig() config.Config {
	var (
		path    = "config/config.json"
		c       = config.Config{}
		content []byte
		err     error
	)

	if content, err = ioutil.ReadFile(path); err != nil {
		log.Fatalf("error: config file %s, %s", path, err.Error())
	}

	if err = json.Unmarshal(content, &c); err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	return c
}
