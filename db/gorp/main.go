package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

func main() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var users []User
	_, _ = dbmap.Select(&users, "select * from users")
	fmt.Println(users)
}

type User struct {
	Id   int
	Name string
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	return dbmap
}
