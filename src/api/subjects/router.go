package subjects

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"models"
	"net/http"
	"store"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getSubjects).Methods("GET")
	router.HandleFunc("/", createSubject).Methods("POST")
	return router
}

func getSubjects(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(store.GetSubjects().ToJson())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func createSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subject models.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err := store.CreateSubject(subject.ToBson()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}
