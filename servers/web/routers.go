package web

import (
	"avito/servers/web/handlers"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/user_banner", handlers.GetUserBanner)
	e.GET("/banners", handlers.GetBanner)
	e.POST("/banners", handlers.CreateBanner)
	e.PATCH("/banner/:id", handlers.UpdateBanner)
	e.DELETE("/banner/:id", handlers.DeleteBanner)
	return e
}
