package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "someone:password@(db:3306)/some_db")
	if err != nil {
		log.Fatal("OpenError: ", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}
}