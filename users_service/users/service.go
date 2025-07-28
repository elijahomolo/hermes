package users

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/elijahomolo/hermes/users_service/db"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Repository interface {
	GetByID(id string) (*User, error)
	Create(user *User) error
	Delete(id string) error
	Update(id string) (*User, error)
	Login(email string) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type Service struct {
	Repo Repository
}

func (s *Service) getUserByEmail(email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	db, err := db.DBConn()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	row := db.QueryRow(query, email)
	user := &User{}
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.Country, &user.Email, &user.CreatedAt)
	if err != nil {
		log.Printf("Error querying user by email: %v", err)
		if err == sql.ErrNoRows {
			log.Printf("No user found with email: %s", email)
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to query user by email: %w", err)
	}
	// Return the user if found
	if user.Email != email {
		log.Printf("User with email %s not found", email)

		return nil, fmt.Errorf("user with email %s not found", email)
	}

	defer db.Close()
	log.Printf("User found: %v", user)

	return user, nil

}

func NewService(r Repository) *Service {
	return &Service{Repo: r}
}

func (s *Service) GetUser(id string) (*User, error) {
	return s.Repo.GetByID(id)
}

func (s *Service) CreateUser(firstName string, lastName string, dateOfBirth string, password string, country string, language string, email string, createdAt string) (*User, error) {
	id := generateID()
	user := &User{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Password:    password, // Placeholder for password handling
		DateOfBirth: dateOfBirth,
		Country:     country,
		Language:    language,
		Email:       email,
		CreatedAt:   createdAt,
	}

	db, err := db.DBConn()
	if err != nil {
		return nil, err
	}

	log.Printf("Creating user: %v", user)

	query := "INSERT INTO users (id, LastName, FirstName, DateOfBirth, Country, Language, Email, Password,  CreatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, user.ID, user.LastName, user.FirstName, user.DateOfBirth, user.Country, user.Language, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		log.Printf("Error inserting user into database: %v", err)
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return user, err

}

func (s *Service) GetUserByEmail(email string) (*User, error) {
	user, err := s.getUserByEmail(email)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *Service) DeleteUser(id string) error {
	return s.Repo.Delete(id)
}

// add update user function

// Placeholder for ID generation logic.
func generateID() string {
	return time.Now().Format("20060102150405")
}

func (s *Service) GenerateID() string {
	return generateID()
}
