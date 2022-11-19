package main

import (
	"connect-postgres/handlers"
	"log"
	"net/http"
)

func (app *Config) routes(l *log.Logger) http.Handler {

	l.Println("in Routes")
	mux := http.NewServeMux()

	rootHandle := handlers.NewRoot(l)

	mux.Handle("/", rootHandle)

	return mux
}
