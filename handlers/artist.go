package handlers

import (
	"net/http"
	"strconv"

	groupie "groupie/data"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}

	artist := groupie.FetchArtist()
	if id > len(artist) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	data.Title = "Artist"
	data.Data = artist[id-1]

	Temp.ExecuteTemplate(w, "base.html", data)
}
