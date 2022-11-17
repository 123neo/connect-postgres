package main

import "net/http"

func (app *Config) routes() http.Handler {
	// hh := handlers.NewRoot(l)
	// gh := handlers.NewHello(l)
	mux := http.NewServeMux()
	// mux.Handle("/", hh)
	// mux.Handle("/hello", gh)

	return mux
}
