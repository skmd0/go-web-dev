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
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);`)
	if err != nil {
		panic(err)
	}
}
