package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/elijahomolo/hermes/users_service/users"

	"github.com/gorilla/mux"
)

var email string
var password string
var userService *users.Service

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement login logic
	// takes input value for name
	fmt.Print("Enter your name: ")
	fmt.Scan(&email)

	if email == " " {
		fmt.Print()
		return
	}

	fmt.Print("Enter your password: ")
	fmt.Scan(&password)
	if password == " " {
		fmt.Print("Password cannot be empty")
		return
	}

	// TODO: Implement user authentication logic here
	_, err := userService.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))

}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Enter your first name: ")
	var firstName string
	fmt.Scan(&firstName)
	if firstName == " " {
		http.Error(w, "First name cannot be empty", http.StatusBadRequest)
		return
	}
	fmt.Print("Enter your last name: ")
	var lastName string
	fmt.Scan(&lastName)
	if lastName == " " {
		http.Error(w, "Last name cannot be empty", http.StatusBadRequest)
		return
	}
	fmt.Print("Enter your date of birth (YYYY-MM-DD): ")
	var dateOfBirth string
	fmt.Scan(&dateOfBirth)
	if dateOfBirth == " " {
		http.Error(w, "Date of birth cannot be empty", http.StatusBadRequest)
		return
	}
	// // Parse dateOfBirth string to time.Time
	// dateOfBirthTime, err := time.Parse("2006-01-02", dateOfBirth)
	// if err != nil {
	// 	http.Error(w, "Invalid date format. Please use YYYY-MM-DD.", http.StatusBadRequest)
	// 	return
	// }
	fmt.Print("Enter your country: ")
	var country string
	fmt.Scan(&country)
	if country == " " {
		http.Error(w, "Country cannot be empty", http.StatusBadRequest)
		return
	}

	var email string
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	if email == " " {
		http.Error(w, "Email cannot be empty", http.StatusBadRequest)
		return
	}
	var password string
	fmt.Print("Enter your password: ")
	fmt.Scan(&password)
	if password == " " {
		http.Error(w, "Password cannot be empty", http.StatusBadRequest)
		return
	}

	var language string
	fmt.Print("Enter your preferred language: ")
	fmt.Scan(&language)
	if language == " " {
		http.Error(w, "Language cannot be empty", http.StatusBadRequest)
		return
	}

	createdAt := time.Now().UTC().Format("2006-01-02 15:04:05")

	user, err := userService.CreateUser(firstName, lastName, dateOfBirth, password, country, language, email, createdAt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created successfully: %v", user)))
}

func main() {
	// Set up HTTP router
	router := mux.NewRouter()

	// Define a simple login handler

	router.HandleFunc("/login", LoginHandler)
	router.HandleFunc("/create_user", CreateUserHandler)
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
