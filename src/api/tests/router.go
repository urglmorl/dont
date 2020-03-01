package tests

import (
	"constants"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(constants.Root, getTests)
	return router
}

func createTest(w http.ResponseWriter, r *http.Request) {
}

func getTests(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("tests"))
	w.WriteHeader(http.StatusOK)
}

func getTestById(w http.ResponseWriter, r *http.Request) {
}

func updateTest(w http.ResponseWriter, r *http.Request) {
}

func deleteTest(w http.ResponseWriter, r *http.Request) {
}
