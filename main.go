package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cfg := mysql.Config{
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DATABASE_ADDRESS"),
		DBName:               os.Getenv("DATABASE_NAME"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to MySQL!")
}
