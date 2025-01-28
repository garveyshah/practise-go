package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// User represents a User in the database
type User struct {
	ID int
	Name string
	Age int
}

func main() {
	// Open the database
	db, err := sql.Open("sqlite3", "practice.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the users table
	if err := createTable(db); err != nil {
		log.Fatal(err)
	}

	if err := insertUser(db, "Alice", 25); err != nil {
		log.Fatal(err)
	}
	if err := insertUser(db, "Bob", 30); err != nil {
		log.Fatal(err)
	}

	// Retrieve all users
	users, err := selectAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users in the database:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	// Update a user's age
	if err := updateUserAge(db, "Alice", 35); err != nil {
		log.Fatal(err)
	}
	log.Println("Alice's age updated")

	// Delete a user
	if err := deleteUser(db, "Bob"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bob deleted!")

	// Final list of users
	users, err = selectAllUsers(db)
	if err != nil {
		log.Fatal(err)
	} 
	fmt.Println("Final users in the database:")
	for _, user := range users {
		fmt.Printf("ID  %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age )
	}
}

// createTable creates the uses table if it doesn't exist
func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	age INTER NOT NULL
	)`
	_, err := db.Exec(query)
	return err
}

// insertUser inserts a new User into the users table
func insertUser(db *sql.DB, name string, age int) error {
	query  := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err := db.Exec(query, name, age)
	return err
}

// selectAllUsers retrieves all users from the users table
func selectAllUsers(db *sql.DB) ([]User, error) {
	query := `SELECT id, name, age FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// updateUserAge updates a user's age by their name
func updateUserAge(db *sql.DB, name string, newAge int) error {
	query := `UPDATE users SET age = ? WHERE name = ?`
	_, err := db.Exec(query, newAge, name)
	return err
}

// deleteUser deletes  a user from the uses table by their name
func deleteUser(db *sql.DB, name string) error {
	query := `DELETE FROM users WHERE name = ?`
	_, err := db.Exec(query, name)
	return err
}