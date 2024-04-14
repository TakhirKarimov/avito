package init

import (
	"avito/db/cache"
	"avito/db/repository"
	"avito/pkg/config"
	"avito/pkg/di"
	"avito/pkg/logger"
	"avito/servers"
	"avito/servers/web"
)

func init() {
	di.AddDefinition(
		&di.Definition{
			Name: "web",
			Build: func(c *di.Container) (interface{}, error) {
				return web.NewWebServer(), nil
			},
		},

		&di.Definition{
			Name: "servers",
			Build: func(c *di.Container) (interface{}, error) {
				serverManager := &servers.ServerManager{
					Servers: make([]servers.Server, 0, 1),
				}
				serverManager.AddServer(di.Get("web").(*web.WebServer))
				return serverManager, nil
			},
		},

		&di.Definition{
			Name: "db",
			Build: func(c *di.Container) (interface{}, error) {
				return repository.NewConnect()
			},
		},
		&di.Definition{
			Name: "cache_db",
			Build: func(c *di.Container) (interface{}, error) {
				return cache.NewCacheConnect()
			},
		},
		&di.Definition{
			Name: "logger",
			Build: func(c *di.Container) (interface{}, error) {
				return logger.NewLogger()
			},
		},

		&di.Definition{
			Name: "config",
			Build: func(c *di.Container) (interface{}, error) {
				return config.NewConfig()
			},
		},
	)
}
