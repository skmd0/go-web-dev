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

	//_, err = db.Exec(`INSERT INTO USERS(age, email, first_name, last_name) VALUES (30, 'domen@skamlic.com', 'Domen', 'Skamlic');`)
	//if err != nil {
	//	panic(err)
	//}

	var id int
	err = db.QueryRow(`SELECT id FROM users WHERE email = $1 AND age > $2;`, "domen@skamlic.com", 16).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

	rows, err := db.Query(`SELECT id FROM users WHERE email = $1 AND age > $2;`, "domen@skamlic.com", 16)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}
