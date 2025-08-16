package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var configDbCmd = &cobra.Command{
	Use:   "config-db",
	Short: "Configure the database connection for the Hermes users service",
	Long:  "This command sets up the database connection for the Hermes users service. It requires environment variables to be set for database connection parameters.",
	Run: func(cmd *cobra.Command, args []string) {

		// Load environment variables from .env file
		if err := godotenv.Load(); err != nil {
			fmt.Errorf("error loading .env file: %w", err)
			return
		}

		cfg := mysql.Config{
			User:                 os.Getenv("DATABASE_USER"),
			Passwd:               os.Getenv("DATABASE_PASSWORD"),
			Net:                  "tcp",
			Addr:                 os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT"),
			AllowNativePasswords: true,
		}

		db, err := checkMySQlConnection(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error connecting to the database: %v\n", err)
			return
		}
		defer db.Close()

		// Create Hermes database if it does not exist

		dbName := os.Getenv("DATABASE_NAME")
		if dbName == "" {
			fmt.Fprintln(os.Stderr, "DATABASE_NAME environment variable is not set")
			return
		}

		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			panic(err)
		}
		defer db.Close()

		_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec("USE " + dbName)
		if err != nil {
			panic(err)
		}

		// Create Users table if it does not exist
		createTableQuery := `
		CREATE TABLE IF NOT EXISTS Users (
			ID VARCHAR(36) PRIMARY KEY,
			FirstName VARCHAR(50) NOT NULL,
			LastName VARCHAR(50) NOT NULL,
			DateOfBirth DATE NOT NULL,
			Password VARCHAR(255) NOT NULL,
			Country VARCHAR(50) NOT NULL,
			Language VARCHAR(50) NOT NULL,
			Email VARCHAR(100) UNIQUE NOT NULL,
			CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
		);`

		if _, err := db.Exec(createTableQuery); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating Users table: %v\n", err)
			return
		}

		fmt.Println("Database and Users table configured successfully.")

	},
}

func checkMySQlConnection(cfg mysql.Config) (*sql.DB, error) {
	fmt.Println("Testing mySQl connection")
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	fmt.Println("Database connection established successfully.")

	return db, nil
}

// Initialize the command

func init() {
	rootCmd.AddCommand(configDbCmd)
}
