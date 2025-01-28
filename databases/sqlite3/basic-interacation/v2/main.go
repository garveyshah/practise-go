package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open the SQLite database (or create it if it doesn't exist)
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the 'users' table if it does't already exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (if INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new user into the 'users' table
	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "Godwin")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted user Godwin into the table")

	// Query the database to confirm the insertion
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the result set and print each row
	if rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("user %d, %s\n", id, name)
	}

	// Check for errors after iterating through the rows
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the updated rows
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("User: %d, %s\n", id, name)
	}

	// Delete a user
	_, err = db.Exec("DELETE FROM users WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted user with ID 1")

	// Final query to see the result
	rows, err = db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the remaining users
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("user: %d, %s\n", id, name)
	}
}
