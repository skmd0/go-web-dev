package main

import (
	"fmt"
	"go-web-dev/mvc/db_service/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "testtest"
	dbName   = "postgres"
)

func main() {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		host, port, user, password, dbName)

	us, err := models.NewUserService(dsn)
	if err != nil {
		panic(err)
	}
	if err := us.DestructiveReset(); err != nil {
		panic(err)
	}
	createFakeUser(us, "Domen Skamlic", "domen@skamlic.com")
	createFakeUser(us, "Domenko Skamlici", "domenko@skamlici.net")
}

func createFakeUser(us *models.UserService, name, email string) {
	u1 := models.User{Name: name, Email: email}
	err := us.Create(&u1)
	if err != nil {
		panic(err)
	}
}
