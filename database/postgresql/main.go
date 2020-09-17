package main

import (
	"database/sql"
	"fmt"
	// this is needed because init() function needs to be called in pq package
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "testtest"
	dbName   = "postgres"
)

func main() {
	var dbConnection string
	dbConnection = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		panic(err)
	}
	// check if connection works
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
}
