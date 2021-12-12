package driver

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type goCache struct {
	cache *cache.Cache
}

func NewGoCache(defaultExpiration, cleanupInterval time.Duration) *goCache {
	return &goCache{cache: cache.New(defaultExpiration, cleanupInterval)}
}

func (g *goCache) Get(key string) (val interface{}, isset bool) {
	return g.cache.Get(key)
}

func (g *goCache) Set(key string, val interface{}, ttl time.Duration) error {
	g.cache.Set(key, val, ttl)
	return nil
}

func (g *goCache) Delete(key string) error {
	g.cache.Delete(key)
	return nil
}

func (g *goCache) Flush() error {
	g.cache.Flush()
	return nil
}
