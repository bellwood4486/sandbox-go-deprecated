package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	type User struct {
		Id   int
		Name string
	}

	db, err := sql.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(`SELECT id, name FROM "users"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	user := User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	fmt.Println(users)

	var r sql.Result
	var count int64
	// insert
	r, err = db.Exec(`INSERT INTO users (name) values ($1)`, "foo")
	if err != nil {
		log.Fatal(err)
	}
	count, _ = r.RowsAffected()
	fmt.Printf("%d inserted!\n", count)

	// update
	r, err = db.Exec(`UPDATE users SET name = $1 WHERE name = $2`, "bar", "foo")
	if err != nil {
		log.Fatal(err)
	}
	count, _ = r.RowsAffected()
	fmt.Printf("%d updated!\n", count)

	// delete
	r, err = db.Exec(`DELETE FROM users WHERE name = $1`, "bar")
	if err != nil {
		log.Fatal(err)
	}
	count, _ = r.RowsAffected()
	fmt.Printf("%d deleted!\n", count)
}
