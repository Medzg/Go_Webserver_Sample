package controllers

import (
	"fmt"
	"net/http"
)

func EventController(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)

	// to do implement crud for the server not just printing messages
	switch request.Method {
	case http.MethodGet: println("this need to be from get")
	case http.MethodPost :println("this need to be from post")
	case http.MethodPut: println("this need to be from put")
	case http.MethodDelete: println("this need to be from delete")
	default:
		println("thee no handler for this")





	}
	fmt.Fprintf(writer,"Welcome From Api")

}
