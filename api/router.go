package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(host string, port string) *echo.Echo {
	e := echo.New()
	e.Any("/", Home)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Version Routing
	buildRoutes(e.Group("/v1"), v1)

	return e
}

func buildRoutes(g *echo.Group, rs routes) {
	for _, route := range rs {
		g.Match(route.Methods, route.PathName, route.HandlerFunc)
	}
}
