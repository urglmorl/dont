package questions

import (
	"constants"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(constants.Root, getQuestions).Methods("GET")
	return router
}

func createQuestion(w http.ResponseWriter, r *http.Request) {
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
}

func getQuestionById(w http.ResponseWriter, r *http.Request) {
}

func updateQuestion(w http.ResponseWriter, r *http.Request) {
}

func deleteQuestion(w http.ResponseWriter, r *http.Request) {
}
