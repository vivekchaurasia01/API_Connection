package main

import (
	"API_CONNECTION/API/internal/users"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type UserData struct{
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
    Email     string `json:"email"`
}

type Server struct{
	UserManager *users.Manager
}

func main(){
	manager := users.NewManager()

	s := Server{
		UserManager: manager,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}",handleRoot)
	mux.HandleFunc("/Goodbye",handleGoodbye)
	mux.HandleFunc("/hello/",handleHelloParameterize)
	mux.HandleFunc("/responses/{user}/hello/",handleUserResponsesHello)
	mux.HandleFunc("POST/user/hello/",s.handleHelloHeader)
	mux.HandleFunc("POST/AddUser",s.handleAddUser)
	mux.HandleFunc("POST/json",handleJSON)

	
	log.Fatal(http.ListenAndServe(":8080",mux))
}

func (s Server)handleHelloHeader(w http.ResponseWriter, r *http.Request){
	firstName := r.Header.Get("userFirst")

	if firstName == "" {
		http.Error(w,"Invalid first name provided",http.StatusBadRequest)
		return 
	}

	lastName := r.Header.Get("userlast")

	if lastName == "" {
		http.Error(w,"Invalid last name provided",http.StatusBadRequest)
		return 
	}

}
func(s Server)handleAddUser(w http.ResponseWriter, r *http.Request){
	ContentType:= r.Header.Get("content-Type") // HTTP has headers + body and Headers describe what the body contains.
	if ContentType!= "application/json"{  //contentType is a string.
		http.Error(w,fmt.Sprintf("unsupported Content Type header: %q",ContentType),http.StatusUnsupportedMediaType)
		return 
	}
	requestBody := http.MaxBytesReader(w,r.Body,1048576) // limit size to prevent Dos and DDos.
	decoder := json.NewDecoder(requestBody)
	decoder.DisallowUnknownFields() //Rejects extra fields not in struct

	var u UserData

	err := decoder.Decode(&u)
	if err != nil{
		slog.Error("error decoding adduser request body","err",err)
		http.Error(w,"bad request body",http.StatusBadRequest)
		return
	}
	err = s.UserManager.AddUser(u.FirstName,u.LastName,u.Email)
		if err != nil {
			http.Error(w,fmt.Sprintf("Error adding user: %v\n",err),http.StatusBadRequest)
			return 
		}
		w.WriteHeader(http.StatusCreated)
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
func handleHelloParameterize(w http.ResponseWriter, r *http.Request){
	params := r.URL.Query()
	username := "User"
	UserList := params["user"]
	if len(UserList) > 0 {
		username = UserList[0]
	}
	handleHello(w,username)
}
func handleUserResponsesHello(w http.ResponseWriter, r *http.Request){
	username := r.PathValue("user")

	handleHello(w,username)
}
func handleJSON(w http.ResponseWriter, r *http.Request){
	byteData,err := io.ReadAll(r.Body)
	if err != nil || len(byteData) < 1 {
		slog.Error("error reading request body","err",err)
		http.Error(w,"bad request body",http.StatusBadRequest)
		return
	}
	// Lets Unmarshall them...
	var reqData UserData
	err = json.Unmarshal(byteData,&reqData)
	if err != nil{
		slog.Error("error Unmarshalling request body","err",err)
		http.Error(w,"error while parsing reqbody JSON",http.StatusBadRequest)
		return 
	}
	if reqData.FirstName == ""{
		http.Error(w,"invalid Username provided",http.StatusBadRequest)
		return 
	}
	handleHello(w,reqData.FirstName)
}
func handleHello(w http.ResponseWriter, username string){
	var output bytes.Buffer
	output.WriteString("hello,")
	output.WriteString(username)
	output.WriteString("!\n")

	_,err := w.Write(output.Bytes())
	if err != nil{
		slog.Error("error writing resdponse body","err",err)
		return
	}
}
