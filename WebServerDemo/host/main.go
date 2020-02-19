package main

import (
	mainMainServer "WebServerDemo/mainServer"
	"fmt"
	"net/http"
)

type  myHandler struct {
	greeting string
}

func  (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte(fmt.Sprintf("%v world",mh.greeting)))
}

/*func  main()  {
	http.Handle("/",&myHandler{greeting: "Hello"})
	http.ListenAndServe(":8000",nil)

}*/
/*func  main()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/" , homeLink)
	fmt.Println(http.ListenAndServe(":8000",router))

}

func homeLink(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
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
*/
func main()  {

 s,error := mainMainServer.NewMainServer("",5000)
 if error != nil{
 	panic(error)
 }
 s.StartMainServer()

}