package user 

import (
"errors"
"time" 
) 

var (
ErrUserNotFound = errors.New("user not found")
)

type Repository interface {
	GetByID(id string) (*User, error)
	Create(user *User) error
	Delete(id string) error
        Update(id string) (*User, error) 
}

type Service struct {
	repo Repository
}


func NewService(r Repository) *Service {
return &Service{repo: r} 
}

func (s *Service) GetUser(id string) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) CreateUser(name, firstName string, lastName string, dateOfBirth time.Time, country string, email string) (*User, error) {
	user := &User{
		ID:        generateID(),
		FirstName: firstName,
		LastName:  lastName,
                DateOfBirth: dateOfBirth, 
                Country:   country,
		Email:     email,
		CreatedAt: time.Now(),
	}
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

// add update user function 

// Placeholder for ID generation logic.
func generateID() string {
	return time.Now().Format("20060102150405")
}


