package api

import (
	"api/auth"
	"api/subjects"
	"api/tests"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	mount(router, "/subjects", subjects.Router())
	mount(router, "/tests", tests.Router())
	mount(router, "/auth", auth.Router())
	return router
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
	w.WriteHeader(http.StatusOK)
}
