package cache

import (
	"time"

	"github.com/google/wire"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var ProviderSet = wire.NewSet(NewOptions, NewGoCache)

type Options struct {
	DefaultExpiration int64
	CleanupInterval   int64
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := &Options{}
	if err = v.UnmarshalKey("cache", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal cache option error")
	}

	return o, err
}

type goCache struct {
	cache *cache.Cache
}

func NewGoCache(o *Options) CacheInterface {
	return &goCache{
		cache: cache.New(time.Duration(o.DefaultExpiration)*time.Second, time.Duration(o.CleanupInterval)),
	}
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
