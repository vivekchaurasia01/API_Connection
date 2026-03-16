package main

import (
	// "API_CONNECTION/API/internal/users"
	"log"
	"log/slog"
	"net/http"
)

// type UserData struct{
// 	FirstName string
// 	LastName string
// 	Email string
// }

// type Server struct{
// 	UserManager *users.Manager
// }

func main(){
	// manager := users.NewManager()

	// s := Server{
	// 	UserManager: manager,
	// }

	mux := http.NewServeMux()
	mux.HandleFunc("/",handleRoot)
	mux.HandleFunc("/Goodbye",handleGoodbye)
	mux.HandleFunc("/hello/",handleHelloParameterize)
	
	log.Fatal(http.ListenAndServe(":8080",mux))
}
func handleRoot(w http.ResponseWriter, _*http.Request){
	_,err := w.Write([]byte("Welcome_User"))
	if err != nil{
		slog.Error("Error while writing Responses","error",err)
	}
}
func handleGoodbye(w http.ResponseWriter, _*http.Request){
	_,err := w.Write([]byte("Goodbye_User(Thank You for using our services)"))
	if err != nil{
		slog.Error("Error while Writing Resposes","error",err)
	}
}
func 
