package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Home!")
}

// v1
func V1(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from v1!")
}

func Animals(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Animals!")
}

func Wildcard(c echo.Context) error {
	req := c.Request()
	url := c.Scheme() + "://" + req.Host + req.RequestURI[:4]
	return c.Redirect(302, url)
}
