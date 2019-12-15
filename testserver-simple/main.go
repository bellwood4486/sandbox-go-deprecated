package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func handle(_ http.ResponseWriter, r *http.Request) {
	fmt.Printf("------\n")
	fmt.Printf("method: %v\n", r.Method)
	fmt.Printf("header: %v\n", r.Header)
	fmt.Printf("uri: %v\n", r.RequestURI)
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read body: %v\n", err)
	}
	fmt.Printf("body: %v\n", string(buf))

	time.Sleep(500 * time.Millisecond)
}

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
