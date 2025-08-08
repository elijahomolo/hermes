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

	router.HandleFunc("/login", users.LoginHandler)
	router.HandleFunc("/create_user", users.CreateUserHandler)
	router.HandleFunc("/get_user_by_email", users.GetUserByEmailHandler)
	router.HandleFunc("/delete_user/{id}", users.DeleteUserHandler)
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
