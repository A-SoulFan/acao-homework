//go:build wireinject
// +build wireinject

package gin

import (
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/server/gin/handler"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/server/gin/middleware"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/server/gin/router"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/service"
	"github.com/A-SoulFan/acao-homework/internal/pkg/cache"
	"github.com/A-SoulFan/acao-homework/internal/pkg/config"
	"github.com/A-SoulFan/acao-homework/internal/pkg/database"
	"github.com/A-SoulFan/acao-homework/internal/pkg/log"
	"github.com/A-SoulFan/acao-homework/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	http.ProviderSet,
	log.ProviderSet,
	database.ProviderSet,
	cache.ProviderSet,

	middleware.NewErrorInterceptor,
	handler.ProviderSet,
	service.ProviderSet,
	router.InitRouter,

	ProviderSet,
)

func InitServer(configFile string) (*Server, error) {
	panic(wire.Build(providerSet))
}
