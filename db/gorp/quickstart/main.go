package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

func main() {
	// initialize the DbMap
	dbmap := initDb()
	defer dbmap.Db.Close()

	// delete any existing rows
	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	// create two posts
	p1 := newPost("Go 1.1 released!", "Lorem ipsum lorem ipsum")
	p2 := newPost("Go 1.2 released!", "Lorem ipsum lorem ipsum")

	// insert rows - auto increment PKs will be set properly after the insert
	err = dbmap.Insert(p1, p2)
	checkErr(err, "Insert failed")

	// use convenience SelectInt
	count, err := dbmap.SelectInt("select count(*) from posts")
	checkErr(err, "select count(*) failed")
	log.Println("Rows after inserting:", count)

	// use convenience SelectStr
	title, err := dbmap.SelectStr(`select "Title" from posts where post_id = $1`, p1.Id)
	checkErr(err, "select title failed")
	log.Println("p1's title: ", title)

	// fetch Post with primary key
	obj, err := dbmap.Get(Post{}, p1.Id)
	checkErr(err, "fetch by primary key failed")
	p3 := obj.(*Post)
	log.Println("p1 = p3:", p1.Id == p3.Id)

	// update a row
	p2.Title = "Go 1.2 is better than ever"
	count, err = dbmap.Update(p2)
	checkErr(err, "Update failed")
	log.Println("Rows updated:", count)

	// fetch one row - note use of "post_id" instead of "Id" since column is aliased
	//
	// Postgres users should use $1 instead of ? placeholders
	err = dbmap.SelectOne(p2, "select * from posts where post_id = $1", p2.Id)
	checkErr(err, "SelectOne failed")
	log.Println("p2 row:", p2)

	// fetch all rows
	var posts []Post
	_, err = dbmap.Select(&posts, "select * from posts order by post_id")
	checkErr(err, "Select failed")
	log.Println("All rows:")
	for x, p := range posts {
		log.Printf("    %d: %v\n", x, p)
	}

	// delete row by PK
	count, err = dbmap.Delete(p1)
	checkErr(err, "Delete failed")
	log.Println("Rows deleted:", count)

	// delete row manually via Exec
	_, err = dbmap.Exec("delete from posts where post_id = $1", p2.Id)
	checkErr(err, "Exec failed")

	// confirm count is zero
	count, err = dbmap.SelectInt("select * from posts")
	checkErr(err, "select count(*) failed")
	log.Println("Row count - should be zero:", count)

	log.Println("Done!")
}

type Post struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`               // Column size set to 50
	Body    string `db:"article_body,size:1024"` // Set both column name and size
}

func newPost(title, body string) *Post {
	return &Post{
		Created: time.Now().UnixNano(),
		Title:   title,
		Body:    body,
	}
}

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
