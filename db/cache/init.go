package cache

import (
	"avito/pkg/config"
	"avito/pkg/di"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type CacheConnect struct {
	Conn *redis.Client
}

func NewCacheConnect() (*CacheConnect, error) {
	cfg := di.Get("config").(*config.Config)
	return &CacheConnect{
		Conn: redis.NewClient(
			&redis.Options{
				Addr:     fmt.Sprintf("%s:%d", cfg.CacheConfig.Address, cfg.CacheConfig.Port),
				Password: "",
				DB:       0,
			}),
	}, nil
}
