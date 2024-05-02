package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

func NewDBConnection() *sql.DB {
	cfg := mysql.Config{
		User:   os.Getenv("DATABASE_USER"),
		Passwd: os.Getenv("DATABASE_PASSWORD"),
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", os.Getenv("DATABASE_ADDRESS"), os.Getenv("DATABASE_PORT")),
		DBName: os.Getenv("DATABASE_NAME"),
		Params: map[string]string{"parseTime": "true"},
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	return db
}
