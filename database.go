package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// SQLDatabase structure
type SQLDatabase struct {
	username     string
	password     string
	ipAddr       string
	port         string
	databaseName string
}

// NewSQLDB function: factory function to create a new relational database instance
func NewSQLDB(username, password, ipAddr, port, databaseName string) *SQLDatabase {
	db := new(SQLDatabase)
	db.username = username
	db.password = password
	db.ipAddr = ipAddr
	db.port = port
	db.databaseName = databaseName
	return db
}

// OpenSQL method: Opens a SQL database, should be started before using the service
func (s *SQLDatabase) OpenSQL() *sql.DB {
	result, err := sql.Open("mysql", s.username+":"+s.password+"@("+s.ipAddr+":"+s.port+")/"+s.databaseName)

	if err != nil {
		log.Fatal("Database is not accessable. Please check your settings!")
	}

	err = result.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return result
}
