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

	artist := groupie.FetchArtist()
	data.Title = "home"
	data.Data = artist

	Temp.ExecuteTemplate(w, "base.html", data)
}
