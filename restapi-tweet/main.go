package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bellwood4486/sandbox-go/restapi-tweet/controllers"
	"github.com/gorilla/mux"

	"github.com/lib/pq"
)

func main() {
	pgURL, err := pq.ParseURL("postgres://puser:ppassword@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	controller := controllers.Controller{}
	router := mux.NewRouter()
	router.HandleFunc("/api/tweets", controller.GetTweets(db)).Methods("GET")
	router.HandleFunc("/api/tweets/{id}", controller.GetTweet(db)).Methods("GET")
	router.HandleFunc("/api/tweets", controller.AddTweet(db)).Methods("POST")
	router.HandleFunc("/api/tweets/{id}", controller.PutTweet(db)).Methods("PUT")
	router.HandleFunc("/api/tweets/{id}", controller.RemoveTweet(db)).Methods("DELETE")
	log.Println("Server up on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
