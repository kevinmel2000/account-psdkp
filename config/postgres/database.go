package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// WriteMysqlDB function for creating database connection for write-access
func WritePostgresDB() *sql.DB {
	return CreateDBConnection(fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("WRITE_DB_HOST"), os.Getenv("WRITE_DB_USER"), os.Getenv("WRITE_DB_PASSWORD"), os.Getenv("WRITE_DB_NAME")))

}

// ReadMysqlDB function for creating database connection for write-access
func ReadPostgresDB() *sql.DB {
	return CreateDBConnection(fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("READ_DB_USER"), os.Getenv("READ_DB_PASSWORD"), os.Getenv("READ_DB_HOST"), os.Getenv("READ_DB_NAME")))

}

// CreateDBConnection function for creating database connection
func CreateDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("postgres", descriptor)
	if err != nil {
		defer db.Close()
		return db
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

// CloseDb function for closing database connection
func CloseDb(db *sql.DB) {
	if db != nil {
		db.Close()
		db = nil
	}
}
