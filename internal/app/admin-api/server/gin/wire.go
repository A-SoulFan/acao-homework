//go:build wireinject
// +build wireinject

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

var providerSet = wire.NewSet(
	config.ProviderSet,
	http.ProviderSet,
	log.ProviderSet,
	database.ProviderSet,
	cache.ProviderSet,
	context.ProviderSet,
	router.InitRouter,

	ProviderSet,
)

func InitServer(configFile string) (*Server, error) {
	panic(wire.Build(providerSet))
}
