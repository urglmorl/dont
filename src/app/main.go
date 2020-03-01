package main

import (
	"api"
	"constants"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"utils"
)

func main() {
	// TODO: Что здесь должно происходить:
	//  Здесь должна быть проверка на существование ключа шифрования сессий и если этого ключа нет - генерировать новый
	//  Можно сохранять ключ шифрования либо в переменные ОС, а можно и в базу данных

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	router := mux.NewRouter().StrictSlash(true)
	utils.Mount(router, constants.RootRoute, api.Router())

	srv := &http.Server{
		Addr:         constants.Port,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	go func() {
		log.Println("Listening on port ", constants.Port)
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
