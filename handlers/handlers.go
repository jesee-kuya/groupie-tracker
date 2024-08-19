package handlers

import (
	"net/http"
	"text/template"

	groupie "groupie/data"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("template/home.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	artist := groupie.FetchArtist()

	temp.Execute(w, artist)
}
