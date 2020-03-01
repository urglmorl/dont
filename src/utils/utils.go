package utils

import (
	"crypto/rand"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func Mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}

func TokenGenerator() (b []byte) {
	b = make([]byte, 32)
	_, _ = rand.Read(b)
	return b
}

func UUIDs() uuid.UUID {
	return uuid.New()
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
