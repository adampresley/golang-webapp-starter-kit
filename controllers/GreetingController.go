package controllers

import (
	"net/http"

	"github.com/adampresley/GoHttpService"
)

/*
HelloWorld returns a JSON object with a message of the classic programmer's
greeting.
*/
func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	GoHttpService.Success(writer, "Hello world!")
}
