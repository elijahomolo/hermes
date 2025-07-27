package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/elijahomolo/hermes/config"
	"github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB, err error) {
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
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("Successfully connected to MySQL!")
	return db, nil
}
