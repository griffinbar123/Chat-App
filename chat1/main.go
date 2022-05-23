package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("my-secret-key"))

func main() {
	serve := http.Server{
		Addr:    "10.0.0.243:8080",
		Handler: context.ClearHandler(http.DefaultServeMux),
	}
	http.HandleFunc("/", Home)
	http.HandleFunc("/signin", SignIn)
	http.HandleFunc("/signup", Index)
	http.HandleFunc("/sign-in", SignI)
	http.HandleFunc("/sign-up", SignO)
	http.HandleFunc("/messages", Messages)
	http.HandleFunc("/ws", WsEndpoint)
	serve.ListenAndServe()
}
