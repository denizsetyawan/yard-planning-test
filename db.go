package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// sesuaikan dengan konfigurasi database
var (
	DBUser     = "postgres"
	DBPassword = "12345"
	DBName     = "yard-planning"
	DBHost     = "localhost"
	DBPort     = "5432"
	SSLMode    = "disable"
)

func InitDB() {
	var err error

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		DBUser, DBPassword, DBName, DBHost, DBPort, SSLMode,
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to open connection:", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping database:", err)
		os.Exit(1)
	}

	fmt.Println("Database connected")
}
