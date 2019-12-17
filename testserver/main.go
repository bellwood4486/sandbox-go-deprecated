package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	addr := flag.String("addr", ":8080", "address to listen")
	sleep := flag.Duration("sleep", 500, "milliseconds to sleep")
	flag.Parse()

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		time.Sleep(*sleep * time.Millisecond)

		buf, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Printf("failed to read body: %v\n", err)
		}
		fmt.Println(string(buf))
		fmt.Println("--------------")
	})

	fmt.Printf("listen to %q\n", *addr)
	fmt.Printf("sleep %d ms\n", *sleep)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
