package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func NewDBConnection() *sql.DB {
	// @todo: fetch config from env file
	cfg := mysql.Config{
		//User:   os.Getenv("DBUSER"),
		User: "root",
		//Passwd: os.Getenv("DBPASS"),
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:33306",
		DBName: "scootin_aboot",
		Params: map[string]string{"parseTime": "true"},
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	return db
}
