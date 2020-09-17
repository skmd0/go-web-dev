package main

import (
	"database/sql"
	"fmt"
	// this is needed because init() function needs to be called in pq package
	_ "github.com/lib/pq"
)

type UserDB struct {
	id        int
	email     string
	firstName string
	lastName  string
	age       int
}

func main() {
	db := setup()
	users := make([]UserDB, 0)
	rows, err := db.Query(`SELECT id, email, age, first_name, last_name FROM users`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user UserDB
		err = rows.Scan(&user.id, &user.email, &user.age, &user.firstName, &user.lastName)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	fmt.Println(users)
}

func setup() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "testtest"
		dbName   = "postgres"
	)

	var dbConnection string
	dbConnection = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		panic(err)
	}
	return db
}
