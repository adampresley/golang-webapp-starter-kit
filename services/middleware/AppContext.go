package middleware

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/context"
)

/*
AppContext holds context data for the application. This can hold information
such as a database connection, session data, user info, and more. Your middlewares
should attach functions to this structure to pass critical data to request
handlers.
*/
type AppContext struct {
	DB *sql.DB
}

/*
StartAppContext is a middleware that should be early in the chain. This
sets up the initial context and attaches important data to the Gorilla
Context which comes across in the request.
*/
func (ctx *AppContext) StartAppContext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		context.Set(request, "db", ctx.DB)

		h.ServeHTTP(writer, request)
	})
}
