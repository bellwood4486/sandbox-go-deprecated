// プログラミング言語Go 第7章のhttp4より

package main

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]float32

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		_, _ = fmt.Fprintf(w, "%s: %f\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	_, _ = fmt.Fprintf(w, "%f\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
