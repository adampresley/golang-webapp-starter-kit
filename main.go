package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/adampresley/golang-webapp-starter-kit/global/logger"
	"github.com/adampresley/golang-webapp-starter-kit/services/listener"
	"github.com/adampresley/golang-webapp-starter-kit/services/middleware"

	"github.com/adampresley/logging"
)

var ip = flag.String("ip", "0.0.0.0", "IP address/hostname to bind this service to")
var port = flag.Int("port", 8080, "Port number to bind this service to")

func main() {
	flag.Parse()

	logger.Log = logging.NewLogger("Webapp Starter Kit")

	setupLogger()
	startService()
}

func setupLogger() {
}

func startService() {
	logger.Log.Info("Starting Webapp Starter Kit")
	shutdownChannel := make(chan bool)
	doneChannel := make(chan os.Signal)

	/*
	 * Start an HTTP listener then wait on a blocking channel for
	 * a terminate signal.
	 */
	startHTTPListener(*ip, *port)

	signal.Notify(doneChannel, syscall.SIGINT, syscall.SIGTERM)
	logger.Log.Info(<-doneChannel)

	logger.Log.Info("Shutting down...")
	close(shutdownChannel)

	logger.Log.Info("Done.")
	os.Exit(0)
}

func startHTTPListener(ip string, port int) {
	appContext := &middleware.AppContext{}

	httpListener := listener.NewHTTPListenerService(ip, port, appContext)

	setupMiddleware(httpListener, appContext)
	setupRoutes(httpListener, appContext)

	go httpListener.StartHTTPListener()
}
