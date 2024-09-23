package handlers

import (
	"net/http"

	groupie "groupie/data"
)

// DateData holds the artist details and their concert dates for rendering.
type DateData struct {
	Details groupie.Artist
	Dt      groupie.Indexxx
}

// DatesHandler handles the HTTP request for displaying concert dates for an artist.
func DatesHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var details DateData
	id := GetId(w, r)
	if id <= 0 {
		return
	}
	details.Details = Artiste[id-1]
	details.Dt = Tarehe.Index[id-1]
	data.Title = "dates"
	data.Data = details
	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
