package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("sqlite3", "/test.db")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer db.Close()

	fmt.Println("SQLite connection successful!")
}
