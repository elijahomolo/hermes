package db

import (
	"database/sql"
	"log"

	"github.com/elijahomolo/hermes/users_service/config"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DBConn() (db *sql.DB, err error) {
	godotenv.Load(".env")
	cfg := mysql.Config{
		User:                 config.GetDatabaseUser(),
		Passwd:               config.GetDatabasePassword(),
		Net:                  "tcp",
		Addr:                 config.GetDatabaseAddress(),
		DBName:               config.GetDatabaseName(),
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// if err = db.Ping(); err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }
	// fmt.Println("Successfully connected to MySQL!")
	return db, nil
}
