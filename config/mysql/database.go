package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func WriteMySqlDB() *sql.DB {
	return CreateDBConnection(fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("WRITE_DB_USER"), os.Getenv("WRITE_DB_PASSWORD"), os.Getenv("WRITE_DB_HOST"), os.Getenv("WRITE_DB_NAME")))

}

// ReadMYSqlDB function for creating database connection for write-access
func ReadMYSqlDB() *sql.DB {
	return CreateDBConnection(fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("READ_DB_USER"), os.Getenv("READ_DB_PASSWORD"), os.Getenv("READ_DB_HOST"), os.Getenv("READ_DB_NAME")))

}

// CreateDBConnection function for creating database connection
func CreateDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("mysql", descriptor)
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
