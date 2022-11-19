package handlers

import (
	"log"
	"net/http"
)

type Root struct {
	l *log.Logger
}

func NewRoot(l *log.Logger) *Root {
	return &Root{l}
}

func (rt *Root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.l.Println("Root Here")
	w.Write([]byte("I am at root"))
}
