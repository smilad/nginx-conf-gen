package routes

import (
	"github.com/labstack/echo/v4"
	controller "nginx/controller/admin/http"
)

func MapAdminHandler(g *echo.Group, handler controller.AdminHandler) {

	// domain routes
	g.POST("/domain", handler.Create())
	g.GET("/domain", handler.GetAll())
	g.DELETE("/domain/:id", handler.Delete())
	// zone routes
	g.POST("/zone", handler.CreateCacheZone())
	g.GET("/zone", handler.GetCacheZone())

}
