package main

import (
	"fmt"
	"net/http"

	handle "groupie/handlers"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handle.Handler)
	fmt.Println("http://localhost:8080/")
	server.ListenAndServe()
}
