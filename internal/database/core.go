package database

import (
	"database/sql"
	"fmt"
	"log"
)

var dbname string = "app.db"

var db *sql.DB

func Connect() *sql.DB {
	fmt.Println("connecting to database")
	db, err := sql.Open("sqlite", dbname)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.Exec("PRAGMA foreign_keys = ON;")

	return db
}

func CreateDatabase() error {
	createUsersTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		role TEXT DEFAULT 'user',
		password TEXT NOT NULL,
		avatar_path TEXT
	);`

	createCoffeeTableSQL := `
	CREATE TABLE IF NOT EXISTS coffees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		brand TEXT,
		country TEXT,
		img_url TEXT,
		tags TEXT,
		description TEXT
	);`

	if _, err := db.Exec(createUsersTableSQL); err != nil {
		return err
	}

	if _, err := db.Exec(createCoffeeTableSQL); err != nil {
		return err
	}

	return nil
}
