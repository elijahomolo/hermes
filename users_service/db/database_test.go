package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// CheckDatabaseExists checks if the specified database exists.
func TestDatabaseExists(t *testing.T) {
	// Load environment variables
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	dbName := "hermes"
	cfg := mysql.Config{
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Check if the database exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?)", dbName).Scan(&exists)
	if err != nil {
		t.Fatalf("Error checking database existence: %v", err)
	}
}

// TestUserTableExists checks if the Users table exists in the database.
func TestUserTableExists(t *testing.T) {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		t.Fatal("DATABASE_NAME environment variable is not set")
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT"),
		DBName:               dbName,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	var users string
	err = db.QueryRow("SELECT table_name FROM information_schema.tables WHERE table_name = 'users';").Scan(&users)
	if err != nil && err != sql.ErrNoRows {
		t.Fatalf("Error checking Users table existence: %v", err)
	}

	if users == "" {
		t.Fatal("Users table does not exist")
	}

	// Optionally, you can check the structure of the Users table
	// Uncomment the following lines if you want to check the structure
	// rows, err := db.Query("DESCRIBE users")
	// if err != nil {
	// 	t.Fatalf("Error describing Users table: %v", err)
	// }
	// defer rows.Close()
	//
	// var columnName, columnType string
	// exists = false
	// for rows.Next() {
	// 	if err := rows.Scan(&columnName, &columnType); err != nil {
	// 		t.Fatalf("Error scanning Users table description: %v", err)
	// 	}
	// 	if columnName == "Email" {
	// 		exists = true
	// 		break
	// 	}
	// }

	// if !exists {
	// 	t.Fatal("Users table does not exist")
	// }
}
