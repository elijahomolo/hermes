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
	ListAllUsers() ([]*User, error)
}

type Service struct {
	Repo Repository
}

func (s *Service) UpdateSuperUser(id int, firstName string, lastName string, email string, password string) (*SuperUser, error) {
	superUser, err := s.getSuperUserByUsernameOrEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to get superadmin: %w", err)
	}

	if firstName != "" {
		superUser.FirstName = firstName
	}
	if lastName != "" {
		superUser.LastName = lastName
	}
	if email != "" {
		superUser.Email = email
	}
	if password != "" {
		superUser.Password = password
	}

	dbConn, err := db.DBConn()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbConn.Close()

	query := "UPDATE superusers SET FirstName = ?, LastName = ?, Email = ?, Password = ? WHERE ID = ?"
	_, err = dbConn.Exec(query, superUser.FirstName, superUser.LastName, superUser.Email, superUser.Password, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update superadmin: %w", err)
	}

	log.Printf("Superadmin with ID %d updated successfully", id)
	return superUser, nil
}

func (s *Service) DeleteSuperUser(id int) error {
	dbConn, err := db.DBConn()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbConn.Close()

	query := "DELETE FROM superusers WHERE ID = ?"
	result, err := dbConn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete superadmin: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("superadmin not found")
	}
	log.Printf("Superadmin with ID %d deleted successfully", id)
	return nil
}

func (s *Service) createSuperUser(SuperUser SuperUser) (*SuperUser, error) {
	dbConn, err := db.DBConn()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbConn.Close()

	query := "INSERT INTO superusers (FirstName, LastName, Email, Password) VALUES (?, ?, ?, ?)"
	_, err = dbConn.Exec(query, SuperUser.FirstName, SuperUser.LastName, SuperUser.Email, SuperUser.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create superadmin: %w", err)
	}

	return &SuperUser, nil
}

func (s *Service) getSuperUserByUsernameOrEmail(email string) (*SuperUser, error) {
	log.Printf("Searching for superadmin with email: %s", email)

	dbConn, err := db.DBConn()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbConn.Close()

	query := "SELECT ID, FirstName, LastName, Email, Password FROM superusers WHERE Email = ? LIMIT 1"

	var superUser SuperUser
	err = dbConn.QueryRow(query, email).Scan(
		&superUser.ID,
		&superUser.FirstName,
		&superUser.LastName,
		&superUser.Email,
		&superUser.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("superadmin not found")
		}
		return nil, fmt.Errorf("failed to query superadmin: %w", err)
	}

	return &superUser, nil
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

func (s *Service) UpdateUser(id string) (*User, error) {
	user, err := s.updateUserByID(id, "", "", "", "", "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return user, nil
}

// UpdateUser updates a user's information.
func (s *Service) updateUserByID(id string, firstName string, lastName string, dateOfBirth string, country string, language string, email string) (*User, error) {
	user, err := s.GetUser(id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if firstName != "" {
		user.FirstName = firstName
	}
	if lastName != "" {
		user.LastName = lastName
	}
	if dateOfBirth != "" {
		user.DateOfBirth = dateOfBirth
	}
	if country != "" {
		user.Country = country
	}
	if language != "" {
		user.Language = language
	}
	if email != "" {
		user.Email = email
	}

	dbConn, err := db.DBConn()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbConn.Close()

	query := "UPDATE Users SET FirstName = ?, LastName = ?, DateOfBirth = ?, Country = ?, Language = ?, Email = ? WHERE ID = ?"
	_, err = dbConn.Exec(query, user.FirstName, user.LastName, user.DateOfBirth, user.Country, user.Language, user.Email, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	log.Printf("User with ID %s updated successfully", id)
	return user, nil
}

func (s *Service) ListAllUsers() ([]*User, error) {
	users, err := s.listAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to list all users: %w", err)
	}
	return users, nil
}

// ListAllUsers retrieves all users from the database.
func (s *Service) listAllUsers() ([]*User, error) {
	dbConn, err := db.DBConn()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbConn.Close()

	query := "SELECT * FROM Users"
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.DateOfBirth, &user.Country, &user.Language, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %w", err)
	}

	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	return users, nil
}
