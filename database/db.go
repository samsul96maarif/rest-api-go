package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "phpmyadmin:localhost@tcp(127.0.0.1:3306)/rest_go_db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
