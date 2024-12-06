package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err := sql.Open("mysql", "user:1234567@tcp(172.20.0.4:3306)/test")

	if err != nil {
		log.Fatal("error db connection:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("error db connection:", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`
	if _, err := DB.Exec(createTableQuery); err != nil {
		log.Fatal("error db table create:", err)
	}

	fmt.Println("DB up")
}

func Close() {
	err := DB.Close()
	if err != nil {
		log.Fatal("error db close:", err)
		return
	}
}
