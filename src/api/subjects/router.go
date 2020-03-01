package subjects

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"models"
	"net/http"
	"store"
	"utils"
)

// router path is "/dont/api/v1/subjects"
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", createSubject).Methods("POST")
	router.HandleFunc("/", getSubjects).Methods("GET")
	router.HandleFunc("/{id}", getSubjectById).Methods("GET")
	router.HandleFunc("/{id}", updateSubject).Methods("PUT")
	router.HandleFunc("/{id}", deleteSubject).Methods("DELETE")
	return router
}

func createSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subject models.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.CheckError(err)
	}
	if subject, err := store.CreateSubject(subject); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.CheckError(err)
	} else {
		err = json.NewEncoder(w).Encode(subject)
		utils.CheckError(err)
	}
}

func getSubjects(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if subjects, err := store.GetSubjects(); err != nil {
		w.WriteHeader(http.StatusNotFound)
		utils.CheckError(err)
	} else {
		err = json.NewEncoder(w).Encode(subjects)
		utils.CheckError(err)
	}
}

func getSubjectById(w http.ResponseWriter, r *http.Request) {

}

func updateSubject(w http.ResponseWriter, r *http.Request) {

}

func deleteSubject(w http.ResponseWriter, r *http.Request) {

}
