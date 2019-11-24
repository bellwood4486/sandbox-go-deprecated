package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	Id   int
	Name string
}

var users = map[int]user{}

func listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s\n", r.Method, r.URL)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s\n", r.Method, r.URL)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "invalid payload: %v\n", err)
		return
	}

	fmt.Printf("%s\n", reqBody)

	var newUser user
	if err := json.Unmarshal(reqBody, &newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "invalid json: %v\n", err)
		return
	}

	users[newUser.Id] = newUser
	w.WriteHeader(http.StatusCreated)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", listUsers).Methods(http.MethodGet)
	r.HandleFunc("/users", createUser).Methods(http.MethodPost)
	return r
}

func main() {
	http.Handle("/", newRouter())
	log.Fatal(http.ListenAndServe(":8080", nil)) // 使い捨てコードなのでFatalで手抜き
}
