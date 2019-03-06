package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	v1 "github.com/nitharios/bastion/router/v1"
)

func New() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Version Routing
	versionOne := e.Group("/v1")
	// TODO: Build out the rest of the v1 routes
	for _, route := range v1.Routes {
		versionOne.Match(route.Methods, route.PathName, route.HandlerFunc)
	}

	return e
}
