package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

type Invoice struct {
	Id       int64
	Created  int64
	Updated  int64
	Memo     string
	PersonId int64 `db:"person_id"`
}

type Person struct {
	Id      int64
	Created int64
	Updated int64
	FName   string
	LName   string
}

type InvoicePersonView struct {
	InvoiceId int64 `db:"invoice_id"`
	PersonId  int64 `db:"person_id"`
	Memo      string
	FName     string
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{LowercaseFields: true}}
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

func main() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	// Create some rows
	p1 := &Person{0, 0, 0, "bob", "smith"}
	err = dbmap.Insert(p1)
	checkErr(err, "Insert failed")

	// notice how we can wire up p1.Id to the invoice easily
	inv1 := &Invoice{0, 0, 0, "xmas order", p1.Id}
	err = dbmap.Insert(inv1)
	checkErr(err, "Insert failed")

	// Run your query
	query := `select i."id" "invoice_id", p."id" "person_id", i."memo", p."fname" 
              from "invoices" i, "people" p
              where i."person_id" = p."id"`

	// pass a slice to Select()
	var list []InvoicePersonView
	_, err = dbmap.Select(&list, query)
	checkErr(err, "select failed")

	// this should test true
	expected := InvoicePersonView{inv1.Id, p1.Id, inv1.Memo, p1.FName}
	if reflect.DeepEqual(list[0], expected) {
		fmt.Println("Woot! My join worked!")
	}
}
