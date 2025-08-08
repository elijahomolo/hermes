package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// LoginHandler handles user login requests.

var email string
var password string

// userService should be initialized properly, e.g. in an init function or via dependency injection
var userService *Service

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

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Print("Enter user ID to delete: ")

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	err := userService.DeleteUser(id)
	if err != nil {
		if err == ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error deleting user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("enter email to search for user: ")
	var email string
	fmt.Scan(&email)

	if email == " " {
		http.Error(w, "Email cannot be empty", http.StatusBadRequest)
		return
	}
	// email := r.URL.Query().Get("email")
	// if email == "" {
	// 	http.Error(w, "Email is required", http.StatusBadRequest)
	// 	return
	// }

	user, err := userService.GetUserByEmail(email)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User found: %v", user)))
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

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Print("Enter user ID to update: ")

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	if _, err := userService.GetUser(id); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Prompt for new user details
	fmt.Println("Enter new user details (leave empty to keep current values):")
	// You can use fmt.Scan or fmt.Scanln to read user input
	// For simplicity, let's assume we are updating firstName, lastName, dateOfBirth, country, language, and email
	// You can also use a struct to hold the user details if needed
	// For example:
	// var userUpdate UserUpdate
	// fmt.Scanln(&userUpdate.FirstName, &userUpdate.LastName, &userUpdate.DateOfBirth, &userUpdate.Country, &userUpdate.Language, &userUpdate.Email)

	// Here we will just prompt for each field individually
	// You can modify this to read from the request body if needed
	// For simplicity, we will just read from the console
	// In a real application, you would typically read from the request body (e.g., JSON)
	// For example, you can use json.NewDecoder(r.Body).Decode(&userUpdate) to read JSON input
	// Here we will just prompt for each field individually
	// You can modify this to read from the request body if needed
	fmt.Print("Enter new first name (leave empty to keep current): ")
	var firstName, lastName, dateOfBirth, country, language, email string
	fmt.Scan(&firstName)
	if firstName == " " {
		firstName = "" // Keep current value
	}

	fmt.Print("Enter new last name (leave empty to keep current): ")
	fmt.Scan(&lastName)
	if lastName == " " {
		lastName = "" // Keep current value
	}

	fmt.Print("Enter new date of birth (YYYY-MM-DD, leave empty to keep current): ")
	fmt.Scan(&dateOfBirth)
	if dateOfBirth == " " {
		dateOfBirth = "" // Keep current value
	}

	fmt.Print("Enter new country (leave empty to keep current): ")
	fmt.Scan(&country)
	if country == " " {
		country = "" // Keep current value
	}

	fmt.Print("Enter new preferred language (leave empty to keep current): ")
	fmt.Scan(&language)
	if language == " " {
		language = "" // Keep current value
	}

	fmt.Print("Enter new email (leave empty to keep current): ")
	fmt.Scan(&email)
	if email == " " {
		email = "" // Keep current value
	}

	// Call the update function in the service
	user, err := userService.UpdateUser(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating user: %v", err), http.StatusInternalServerError)
		return
	}

	// Here you can return a success response or updated user details
	// For example, you can return a JSON response with the updated user details
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)

	// For simplicity, we will just print a success message
	fmt.Fprintf(w, "User updated successfully: ID=%s, FirstName=%s, LastName=%s, DateOfBirth=%s, Country=%s, Language=%s, Email=%s",
		user.ID, user.FirstName, user.LastName, user.DateOfBirth, user.Country, user.Language, user.Email)

	// // Optionally, you can return a JSON response with the updated user details
	// // w.Header().Set("Content-Type", "application/json")
	// // json.NewEncoder(w).Encode(user)
	// if language != "" {
	// 	user.Language = language
	// }
	// if email != "" {
	// 	user.Email = email
	// }

	// err = userService.UpdateUserByID(id, firstName, lastName, dateOfBirth, country, language, email)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error updating user: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}
