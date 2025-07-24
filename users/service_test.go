package user

import (
	"errors"
	"testing"
	"time"
)

// MockRepository is a mock implementation of Repository.
type MockRepository struct {
	users map[string]*User
}

func NewMockRepository() *MockRepository {
	return &MockRepository{users: make(map[string]*User)}
}

func (m *MockRepository) GetByID(id string) (*User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (m *MockRepository) Create(user *User) error {
	if _, exists := m.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	m.users[user.ID] = user
	return nil
}

func (m *MockRepository) Delete(id string) error {
	if _, exists := m.users[id]; !exists {
		return ErrUserNotFound
	}
	delete(m.users, id)
	return nil
}

func (m *MockRepository) Update(id string) (*User, error) {
         user, ok := m.users[id]
         if !ok {
              return nil, ErrUserNotFound
            }
         if _, exists := m.users[id]; exists {
                return nil, errors.New("user already exists")
        }
        
        return user, nil
}

// Unit Tests

func TestCreateUser(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	user, err := service.CreateUser("Alice", "Doe", time.Now() , "USA", "alice@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.FirstName != "Alice" || user.Email != "alice@example.com" {
		t.Errorf("user fields not set correctly")
	}

	if user.CreatedAt.IsZero() {
		t.Errorf("expected CreatedAt to be set")
	}
}

func TestGetUser(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	// Create a user manually in the mock repo
	user := &User{
		ID:        "12345",
		FirstName:      "Bob",
		Email:     "bob@example.com",
		CreatedAt: time.Now(),
	}
	repo.Create(user)

	// Fetch user
	got, err := service.GetUser("12345")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.FirstName != "Bob" {
		t.Errorf("expected name 'Bob', got %s", got.FirstName)
	}
}

func TestDeleteUser(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	user := &User{
		ID:        "67890",
		FirstName:      "Charlie",
		Email:     "charlie@example.com",
		CreatedAt: time.Now(),
	}
	repo.Create(user)

	err := service.DeleteUser("67890")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = service.GetUser("67890")
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}
}
