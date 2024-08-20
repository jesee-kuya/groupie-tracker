package main

import (
	"net/http"

	handle "groupie/handlers"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", handle.HomeHandler)
	http.HandleFunc("/artist", handle.ArtistHandler)
	server.ListenAndServe()
}
