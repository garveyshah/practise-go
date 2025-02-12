package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./../db/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert a new user info into the table
	promp := `
	CREATE TABLE user (
	name TEXT NOT NULL,
	email TEXT UNITQUE NOT NULL,
	password Text NOT NULL,
	role TEXT NOT NULL CHECK(role IN ('admin', 'adopter))
	)
	`
	_, err = db.Exec(promp)
	if err != nil {
		log.Fatal(err)
	}
}