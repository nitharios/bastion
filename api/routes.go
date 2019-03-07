package api

import (
	"github.com/labstack/echo/v4"
)

type Route struct {
	Methods     []string
	PathName    string
	HandlerFunc func(echo.Context) error
}

type routes []Route

/********** Version 1 **********/
var v1 = routes{
	Route{
		[]string{"GET"},
		"/",
		V1,
	},
	Route{
		[]string{"GET", "POST"},
		"/animals",
		Animals,
	},
	Route{
		[]string{"GET"},
		"/*",
		Wildcard,
	},
}
