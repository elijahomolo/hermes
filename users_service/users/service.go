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

func (s *Service) deleteUser(id string) error {
	dbConn, err := db.DBConn()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbConn.Close()

	query := "DELETE FROM Users WHERE ID = ?"
	result, err := dbConn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrUserNotFound
	}
	log.Printf("User with ID %s deleted successfully", id)
	return nil
}

// getUserByEmail retrieves a user by their email address.

func (s *Service) getUserByEmail(email string) (*User, error) {
	query := "SELECT * FROM Users WHERE Email = ?"
	dbConn, err := db.DBConn()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbConn.Close()

	row := dbConn.QueryRow(query, email)
	log.Print(row)
	user := &User{}
	err = row.Scan(&user.ID, &user.LastName, &user.FirstName, &user.DateOfBirth, &user.Country, &user.Language, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
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
	return s.deleteUser(id)
}

// add update user function

// Placeholder for ID generation logic.
func generateID() string {
	return time.Now().Format("20060102150405")
}

func (s *Service) GenerateID() string {
	return generateID()
}
