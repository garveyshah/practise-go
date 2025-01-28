package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open the database
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the users table if it doesn't exist
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	age INTEGER NOT NULL
	)
	`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created successfully!")

	// Insert sample data
	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?,?)", "Bob", 30)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sample data inserted successfully!")

	// Query the data
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	log.Println("Users in the database:")
	for rows.Next() {
		var id int
		var name string
		var age int

		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// Update a user's page
	_, err = db.Exec("UPDATE users SET age = ? WHERE name = ?", 35, "Alice")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Alice's age updated successfully!")

	// Delete a user
	_, err = db.Exec("DELETE FROM users WHERE name = ?", "Bob")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bob deleted from the database!")
}
