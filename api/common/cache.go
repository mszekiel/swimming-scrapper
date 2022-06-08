package common

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

var Cache *bigcache.BigCache

func InitCache() {
	_cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))

	Cache = _cache

}
