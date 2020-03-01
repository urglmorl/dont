package api

import (
	"api/auth"
	"api/questions"
	"api/subjects"
	"api/tests"
	"api/themes"
	"constants"
	"github.com/gorilla/mux"
	"net/http"
	"utils"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(constants.Root, index)
	utils.Mount(router, constants.SubjectRoute, subjects.Router())
	utils.Mount(router, constants.ThemeRoute, themes.Router())
	utils.Mount(router, constants.QuestionRoute, questions.Router())
	utils.Mount(router, constants.TestRoute, tests.Router())
	utils.Mount(router, constants.AuthRoute, auth.Router())
	return router
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
	w.WriteHeader(http.StatusOK)
}
