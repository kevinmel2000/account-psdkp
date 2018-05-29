package mssql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var b2cAuthDB, b2cLegacyDB *sql.DB

// GetAuthMsSQLDB function for creating database connection from Auth MSSQL Database
func GetAuthMsSQLDB() *sql.DB {
	if b2cAuthDB == nil {
		b2cAuthDB = CreateDBConnection(os.Getenv("MSSQL_IDENTITY_CONNECTION_STRING"))
	}

	return b2cAuthDB
}

// GetLegacyMsSQLDB function for creating database connection from Legacy MSSQL Database
func GetLegacyMsSQLDB() *sql.DB {
	var port int
	port = 1433

	if b2cLegacyDB == nil {
		b2cLegacyDB = CreateDBConnection(fmt.Sprintf("servers=%s;user id=%s;password=%s;port=%d;database=%s",
			os.Getenv("LEGACY_DB_HOST"), os.Getenv("LEGACY_DB_USER"), os.Getenv("LEGACY_DB_PASSWORD"), port, os.Getenv("LEGACY_DB_NAME")))
	}
	return b2cLegacyDB
}

// CreateDBConnection function for creating  database connection
func CreateDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("mssql", descriptor)
	if err != nil {
		defer db.Close()
		return db
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

// CloseDb function for closing database connection
func CloseDb() {
	if b2cAuthDB != nil {
		b2cAuthDB.Close()
		b2cAuthDB = nil
	}

	if b2cLegacyDB != nil {
		b2cLegacyDB.Close()
		b2cLegacyDB = nil
	}
}
