package main

import (
	"api"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sessions"
	"store"
	"strings"
	"time"
)

func main() {
	// TODO: Что здесь должно происходить:
	//  Здесь должна быть проверка на существование ключа шифрования сессий и если этого ключа нет - генерировать новый
	//  Можно сохранять ключ шифрования либо в переменные ОС, а можно и в базу данных

	sessions.Init()
	store.Init()

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	host := "/dont/api/v1/"

	router := mux.NewRouter().StrictSlash(true)
	mount(router, host, api.Router())

	srv := &http.Server{
		Addr:         ":32678",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	go func() {
		log.Println("Listening on port ", ":32678")
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err := srv.Shutdown(ctx)
	checkErr(err)
	defer cancel()
	log.Println("Server gracefully stopped!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
