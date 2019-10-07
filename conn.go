package main

import "database/sql"

var pool *sql.DB

type connCreds struct {
	Driver       string
	User         string
	Password     string
	Port         int
	DatabaseName string
}

func newDatabase(config string) *sql.DB {
	var connection connCreds

	db, err := sql.Open(connection.Driver, connection.User+":"+connection.Password+"@/"+connection.DatabaseName)

	if err != nil {
		// this really shouldn't happen
		panic(err)
	}

	return db
}

func dbQuery() {
	// code goes here
}
