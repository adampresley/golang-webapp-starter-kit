package main

import (
	"github.com/adampresley/golang-webapp-starter-kit/controllers"
	"github.com/adampresley/golang-webapp-starter-kit/services/listener"
	"github.com/adampresley/golang-webapp-starter-kit/services/middleware"
)

/*
Add routes here using AddRoute and AddRouteWithMiddleware.
*/

func setupRoutes(httpListener *listener.HTTPListenerService, appContext *middleware.AppContext) {
	httpListener.
		AddStaticRoute("/assets/", "./www/assets").
		AddRoute("/greeting", controllers.HelloWorld, "GET")
}
