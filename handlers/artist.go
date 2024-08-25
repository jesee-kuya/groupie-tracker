package handlers

import (
	"net/http"
	"strconv"

	groupie "groupie/data"
)

type Data struct {
	Artist   groupie.Artist
	Relation groupie.Indexxxx
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var relate Data
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

	artist, err := groupie.FetchArtist()
	relations, err1 := groupie.FetchRelation()
	if err != nil || err1 != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
	if id > len(artist) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	relate.Artist = artist[id-1]
	relate.Relation = relations.Index[id-1]
	data.Title = "Artist"
	data.Data = relate

	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
