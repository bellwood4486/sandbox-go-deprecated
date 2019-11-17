package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	type User struct {
		Id   int
		Name string
	}

	db, err := gorm.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var users []User
	db.Find(&users) // SELECT * FROM users;
	fmt.Println(users)
}
