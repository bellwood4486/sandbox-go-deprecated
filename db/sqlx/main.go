package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	type User struct {
		Id   int
		Name string
	}

	db, err := sqlx.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	err = db.Select(&users, "select id, name from users")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}
