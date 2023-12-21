package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func migrate(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
		age integer,
		gender varchar(10),
        password TEXT NOT NULL,
		is_verified integer,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
	
	CREATE TABLE IF NOT EXISTS swipes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		swipe_user_id INTEGER NOT NULL,
		is_swipe_user_id integer NOT NULL,
		swipe_type TEXT CHECK(swipe_type IN ('Like', 'Pass')),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (swipe_user_id) REFERENCES users(id),
		FOREIGN KEY (is_swipe_user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS packages (
		id INTEGER PRIMARY KEY,
		subscription_type TEXT,
		subscription_long integer,
		price REAL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	INSERT INTO packages (subscription_type, subscription_long, price) VALUES ('Gold', 1, 10000.00);
	INSERT INTO packages (subscription_type, subscription_long, price) VALUES ('Silver', 2, 20000.00);
	INSERT INTO packages (subscription_type, subscription_long, price) VALUES ('Bronze', 3,30000.00);

	CREATE TABLE IF NOT EXISTS subscriptions (
        id INTEGER PRIMARY KEY,
        user_id INTEGER,
        packages_id,
        StartDate date,
        EndDate date,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id)
        FOREIGN KEY (packages_id) REFERENCES packages(id)
    );

	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Ganti dengan jalur ke file database SQLite Anda
	db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Lakukan migrasi
	migrate(db)
}
