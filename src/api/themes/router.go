package themes

import (
	"constants"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(constants.Root, getThemes).Methods("GET")
	return router
}

func createTheme(w http.ResponseWriter, r *http.Request) {
}

func getThemes(w http.ResponseWriter, r *http.Request) {
}

func getThemeById(w http.ResponseWriter, r *http.Request) {
}

func updateTheme(w http.ResponseWriter, r *http.Request) {
}

func deleteTheme(w http.ResponseWriter, r *http.Request) {
}
