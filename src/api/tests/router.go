package tests

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getTests)
	return router
}

func getTests(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("tests"))
	w.WriteHeader(http.StatusOK)
}
