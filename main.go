package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// Database configuration for MySQL
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "mypassword",
		Addr:                 "",
		DBName:               "projectmanager",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Create a MySQL storage instance
	sqlStorage := NewMySQLStorage(cfg)

	// Connect to the database
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Create a store instance using the database connection
	store := NewStore(db)

	// Create an API server instance with the store
	api := NewAPIServer(":3000", store)

	// Start the API server
	api.Serve()
}
