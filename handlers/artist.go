package handlers

import (
	"net/http"

	groupie "groupie/data"
)

type Data struct {
	Artist   groupie.Artist
	Relation groupie.Indexxxx
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var relate Data

	id := GetId(w, r)
	if id <= 0 {
		return
	}

	relate.Artist = Artiste[id-1]
	relate.Relation = Rapports.Index[id-1]
	data.Title = "Artist"
	data.Data = relate

	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
