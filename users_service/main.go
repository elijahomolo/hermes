package main

import (
	"log"
	"net/http"

	"github.com/elijahomolo/hermes/users_service/cmd"
	"github.com/elijahomolo/hermes/users_service/users"

	"github.com/gorilla/mux"
)

func main() {
	cmd.Execute()
	// Set up HTTP router
	router := mux.NewRouter()

	// Define a simple login handler

	router.HandleFunc("/login", users.LoginHandler).Methods("POST")
	router.HandleFunc("/create_user", users.CreateUserHandler).Methods("POST")
	router.HandleFunc("/get_user_by_email", users.GetUserByEmailHandler).Methods("POST")
	router.HandleFunc("/delete_user/{id}", users.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/superadmin/login", users.SuperUserLoginHandler).Methods("POST")
	router.HandleFunc("/superadmin/create", users.CreateSuperUserHandler).Methods("POST")
	router.HandleFunc("/superadmin/delete", users.DeleteSuperUserHandler).Methods("POST")
	// router.HandleFunc("/users", userService.CreateUserHandler).Methods("POST")
	// router.HandleFunc("/users/{id}", userService.GetUserHandler).Methods("GET")
	// router.HandleFunc("/users/{id}", userService.UpdateUserHandler).Methods("PUT")
	// router.HandleFunc("/users/{id}", userService.DeleteUserHandler).Methods("DELETE")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
