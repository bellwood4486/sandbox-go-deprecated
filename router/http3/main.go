// プログラミング言語Go 第7章のhttp3より

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

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
