package controllers

import (
	"net/http"

	"github.com/adampresley/golang-webapp-starter-kit/model/response"

	"github.com/adampresley/GoHttpService"
)

/*
HelloWorld returns a JSON object with a message of the classic programmer's
greeting.
*/
func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	greetingResponse := response.GreetingResponse{
		Message: "Hello world!",
	}

	GoHttpService.WriteJson(writer, greetingResponse, 200)
}
