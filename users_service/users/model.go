package users

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	Country     string `json:"country"`
	Language    string `json:"language"`
	Email       string `json:"email"`
	CreatedAt   string `json:"created_at"`
}

type SuperUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
