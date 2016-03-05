package main

import (
	"flag"
	"os"

	"github.com/adampresley/golang-webapp-starter-kit/middleware"
	"github.com/adampresley/golang-webapp-starter-kit/server"
	"github.com/adampresley/sigint"

	"github.com/adampresley/logging"
)

func main() {
	flag.Parse()
	log := logging.NewLoggerWithMinimumLevel("Webapp Starter Kit", logging.StringToLogType(*logLevel))

	if *enableColor {
		log.EnableColors()
	}

	appContext := &middleware.AppContext{
		Log: log,
	}

	httpListener := server.NewHTTPListenerService(*ip, *port, appContext)

	setupMiddleware(httpListener, appContext)
	setupRoutes(httpListener, appContext)

	go func() {
		err = httpListener.StartHTTPListener()
		if err != nil {
			log.Fatalf("Error starting HTTP server: %s")
			os.Exit(1)
		}
	}()

	/*
	 * Block this thread until we receive SIGINT or
	 * SIGTERM
	 */
	doneChannel := make(chan bool)

	sigint.Listen(func() {
		log.Info("Shutting down...")
		doneChannel <- true
	})

	<-doneChannel
	log.Info("Done.")
}
