package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

type Invoice struct {
	Id       int64  `db:"id"`
	Created  int64  `db:"created"`
	Updated  int64  `db:"updated"`
	Memo     string `db:"memo"`
	PersonId int64  `db:"person_id"`
}

type Person struct {
	Id      int64  `db:"id"`
	Created int64  `db:"created"`
	Updated int64  `db:"updated"`
	FName   string `db:"fname"`
	LName   string `db:"lname"`
}

func main() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	p1 := &Person{0, 0, 0, "bob", "smith"}
	inv1 := &Invoice{0, 0, 0, "xmas order", p1.Id}

	err = insertInv(dbmap, inv1, p1)
	checkErr(err, "insertInv failed")

	log.Println("Done!")
}

func insertInv(dbmap *gorp.DbMap, inv *Invoice, per *Person) error {
	// Start a new transaction
	trans, err := dbmap.Begin()
	if err != nil {
		return err
	}

	err = trans.Insert(per)
	checkErr(err, "Insert failed")

	inv.PersonId = per.Id
	err = trans.Insert(inv)
	checkErr(err, "Insert failed")

	// if the commit is successful, a nil error is returned
	return trans.Commit()
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Invoice{}, "invoices").SetKeys(true, "Id")
	dbmap.AddTableWithName(Person{}, "people").SetKeys(true, "Id")

	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
