package main

import (
	"github.com/adampresley/golang-webapp-starter-kit/controllers"
	"github.com/adampresley/golang-webapp-starter-kit/services/listener"
)

/*
Add routes here using AddRoute and AddRouteWithMiddleware.
*/

func setupRoutes(httpListener *listener.HTTPListenerService) {
	httpListener.
		AddRoute("/greeting", controllers.HelloWorld, "GET")
}
