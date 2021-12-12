package cache

import "time"

type Interface interface {
	Get(key string) (val interface{}, isset bool)
	Set(key string, val interface{}, ttl time.Duration) error
	Delete(key string) error
	Flush() error
}
