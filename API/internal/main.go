package main

import "API_CONNECTION/API/internal/users"

type UserData struct{
	FirstName string
	LastName string
	Email string
}

type Server struct{
	UserManager *users.Manager
}
func main(){
	manager := users.NewManager()

	s := Server{
		UserManager: manager,
	}

	

}