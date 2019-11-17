package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

type Person struct {
	Id int64
}

func (p *Person) PreDelete(s gorp.SqlExecutor) error {
	query := `delete from "invoices" where "person_id" = $1`
	r, err := s.Exec(query, p.Id)
	if err != nil {
		return err
	}

	count, _ := r.RowsAffected()
	log.Println("Deleted invoice rows: ", count)

	return nil
}

type Invoice struct {
	Id       int64
	Created  int64
	Updated  int64
	Memo     string
	PersonId int64 `db:"person_id"`
}

func (i *Invoice) PreInsert(s gorp.SqlExecutor) error {
	i.Created = time.Now().UnixNano()
	i.Updated = i.Created
	return nil
}

func (i *Invoice) PreUpdate(s gorp.SqlExecutor) error {
	i.Updated = time.Now().UnixNano()
	return nil
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

	p1 := &Person{0}
	err = dbmap.Insert(p1)
	checkErr(err, "Insert failed")

	inv1 := &Invoice{0, 0, 0, "xmas order", p1.Id}
	err = dbmap.Insert(inv1)
	checkErr(err, "Insert failed")

	inv1.Memo = "new year order"
	_, err = dbmap.Update(inv1)
	checkErr(err, "Update failed")

	count, err := dbmap.Delete(p1)
	checkErr(err, "Delete failed")
	log.Println("Deleted person rows: ", count)

	log.Println("Done!")
}
