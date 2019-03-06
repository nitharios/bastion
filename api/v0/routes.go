package v0

import "net/http"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var Router = Routes{
	Route{
		"Home",
		"GET",
		"/",
		Home,
	},
}
