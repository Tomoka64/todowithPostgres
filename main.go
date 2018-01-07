package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)
	//defined in route_main.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	//defined in route_todo.go
	mux.HandleFunc("/todo/new", newTodo)
	mux.HandleFunc("/todo/create", createTodo)
	mux.HandleFunc("todo/show", showTodo)

	http.ListenAndServe(":8080", nil)
}
