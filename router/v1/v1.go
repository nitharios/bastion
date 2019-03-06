package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Route struct {
	Methods     []string
	PathName    string
	HandlerFunc func(echo.Context) error
}

type routes []Route

var Routes = routes{
	Route{
		[]string{"GET"},
		"/",
		Home,
	},
	Route{
		[]string{"GET", "POST"},
		"/animals",
		Animals,
	},
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func New() string {
	return "Hello World from v1"
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Home!")
}

func Animals(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Animals!")
}
