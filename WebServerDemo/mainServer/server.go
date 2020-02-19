package ServerPackage

import (
	"WebServerDemo/controllers"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MainServer struct {

	url string

}

func  NewMainServer(url string , port int64) (*MainServer,error) {
	if port != 0 {
		if len(url) == 0 {
			return &MainServer{url: fmt.Sprintf(":%d", port)}, nil
		}else{
			return &MainServer{url: fmt.Sprintf("%s:%d",url,port)},nil
		}
	}else {
		return  &MainServer{}, errors.New("port need to be specified")
	}
}

func (s *MainServer) StartMainServer(){

	fmt.Println(fmt.Sprintf("Starting server on :  %s .....",s.url))

	handler := setupControllers()
	http.ListenAndServe(s.url,handler)



}
func setupControllers() *mux.Router  {

	fmt.Println("Setting-up controllers")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/" , controllers.EventController)
	return  router
}



