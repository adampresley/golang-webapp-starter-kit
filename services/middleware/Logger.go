package middleware

import (
	"net/http"
	"time"

	"github.com/adampresley/golang-webapp-starter-kit/global/logger"
)

/*
Logger is a middleware which logs requests to the console. It also includes the
time it takes for the request to complete.
*/
func (ctx *AppContext) Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(writer, request)
		logger.Log.Info(request.Method, "-", request.URL.String(), "(", time.Since(startTime), ")")
	})
}
