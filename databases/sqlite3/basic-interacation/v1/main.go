package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// Open the SQLite database (or create it if it doesn't exist)
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the 'users' table if it doesn't already exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
		if err != nil {
			log.Fatal(err)
		}

		// Insert a new user into the 'user' table
		_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Inserted user Alice into the table")

		// Query the database to confirm the insertion
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Iterate through the result set and print each row
		for rows.Next() {
			var id int
			var name string
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("User: %d, %s\n", id, name)
		}

		// Delete a user by their ID
		_, err = db.Exec("DELETE FROM users WHERE id = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Deleted user with ID 1")

		// Check for errorss after each iteration thruogh the rows
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
}