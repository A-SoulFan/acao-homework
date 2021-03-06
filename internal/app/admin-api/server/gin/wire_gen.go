// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package gin

import (
	"github.com/A-SoulFan/acao-homework/internal/app/admin-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/admin-api/server/gin/router"
	"github.com/A-SoulFan/acao-homework/internal/pkg/cache"
	"github.com/A-SoulFan/acao-homework/internal/pkg/config"
	"github.com/A-SoulFan/acao-homework/internal/pkg/database"
	"github.com/A-SoulFan/acao-homework/internal/pkg/log"
	"github.com/A-SoulFan/acao-homework/internal/pkg/transports/http"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitServer(configFile string) (*Server, error) {
	viper, err := config.New(configFile)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.NewLogger(options)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := database.NewDatabase(databaseOptions)
	if err != nil {
		return nil, err
	}
	cacheOptions, err := cache.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	cacheInterface := cache.NewGoCache(cacheOptions)
	serviceContext := context.NewServiceContext(logger, db, cacheInterface)
	initRouters := router.InitRouter(serviceContext, logger)
	engine := http.NewRouter(httpOptions, logger, initRouters)
	server, err := http.NewServer(httpOptions, logger, engine)
	if err != nil {
		return nil, err
	}
	ginServer := NewServer(logger, server)
	return ginServer, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, http.ProviderSet, log.ProviderSet, database.ProviderSet, cache.ProviderSet, context.ProviderSet, router.InitRouter, ProviderSet)
