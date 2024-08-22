package handlers

import (
	"net/http"

	groupie "groupie/data"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	artist, err := groupie.FetchArtist()
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
	data.Title = "home"
	data.Data = artist

	Temp.ExecuteTemplate(w, "base.html", data)
}
