package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

// NewMySQLStorage creates a new MySQLStorage instance and establishes a connection to the database
func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN()) // Opens a connection to the database using the provided configuration
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping() // Attempts to establish a connection to the database to verify connectivity

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL")

	return &MySQLStorage{db: db} // Returns a pointer to the newly created MySQLStorage instance

}

// Init returns the existing database connection
func (s *MySQLStorage) Init() (*sql.DB, error) {
	return s.db, nil // Returns the database connection stored in the MySQLStorage struct
}
