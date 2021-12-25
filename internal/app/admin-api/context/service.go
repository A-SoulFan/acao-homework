package context

import (
	"context"
	"github.com/A-SoulFan/acao-homework/internal/pkg/cache"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewServiceContext)

type ServiceContext struct {
	db     *gorm.DB
	logger *zap.Logger
	cache  cache.CacheInterface
}

func NewServiceContext(logger *zap.Logger, db *gorm.DB, cache cache.CacheInterface) *ServiceContext {
	return &ServiceContext{
		db:     db,
		logger: logger,
		cache:  cache,
	}
}

func (svc *ServiceContext) GetCache() cache.CacheInterface {
	return svc.cache
}

func (svc *ServiceContext) WithDatabaseContext(ctx context.Context) *gorm.DB {
	return svc.db.WithContext(ctx)
}
