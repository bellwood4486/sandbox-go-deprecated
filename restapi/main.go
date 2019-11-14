// See: https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type structError struct {
	Kind    string
	Message string
}

type event struct {
	ID          string
	Title       string
	Description string
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
	},
}

func homeLink(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "invalid payload: %s", err)
		return
	}

	var newEvent event
	if err := json.Unmarshal(reqBody, &newEvent); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "invalid json: %s", err)
		return
	}
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newEvent); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "json encoding failed: %s", err)
		return
	}
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	for _, e := range events {
		if e.ID == eventId {
			setOk(w, e)
			return
		}
	}
	setNotFound(w, eventId)
}

func getAllEvents(w http.ResponseWriter, _ *http.Request) {
	_ = json.NewEncoder(w).Encode(events)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "invalid payload: %s", err)
		return
	}

	var updatedEvent event
	if err := json.Unmarshal(reqBody, &updatedEvent); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "invalid json: %s", err)
		return
	}
	for i := range events {
		e := &events[i]
		if e.ID == eventId {
			e.Title = updatedEvent.Title
			e.Description = updatedEvent.Description
			_ = json.NewEncoder(w).Encode(e)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventId, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "id unspecified")
	}
	for i := range events {
		if events[i].ID == eventId {
			events = append(events[:i], events[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprintf(w, "event not found: %q", eventId)
}

func setOk(w http.ResponseWriter, v interface{}) {
	setJSONResponseHeader(w, http.StatusOK)
	setJSONResponseBody(w, v)
}

func setNotFound(w http.ResponseWriter, id string) {
	setJSONResponseHeader(w, http.StatusNotFound)
	setJSONResponseBody(w, structError{"not_found", id})
}

func setInternalServerError(w http.ResponseWriter, err error) {
	setJSONResponseHeader(w, http.StatusInternalServerError)
	// InternalServerErrorでのエンコード失敗は、それ以上救う手立てがないのでpanicにする。
	if err := json.NewEncoder(w).Encode(structError{"server_error", fmt.Sprint(err)}); err != nil {
		panic(err)
	}
}

func setJSONResponseHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
}

func setJSONResponseBody(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		setInternalServerError(w, err)
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeLink)
	r.HandleFunc("/event", createEvent).Methods("POST")
	r.HandleFunc("/events", getAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	r.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	r.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	return r
}

func main() {
	http.Handle("/", newRouter())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
