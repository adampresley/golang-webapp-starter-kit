package middleware

import (
	"net/http"

	"github.com/adampresley/golang-webapp-starter-kit/global/logger"
)

/*
Authorization is a middleware for handling authorization of requests. Here is
where you might authorize a token, user session, or cookie.
*/
func (ctx *AppContext) Authorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logger.Log.Info("In authorization middleware")
		h.ServeHTTP(writer, request)
	})
}
