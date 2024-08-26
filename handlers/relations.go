package handlers

import (
	"net/http"
	"strconv"

	groupie "groupie/data"
)

type RelData struct {
	Details groupie.Artist
	Rel     groupie.Indexxxx
}

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var detail RelData
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

	relations, err := groupie.FetchRelation()
	artist, err1 := groupie.FetchArtist()

	if id >= len(artist) || id >= len(relations.Index) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	if err != nil || err1 != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
	if id > len(relations.Index) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	detail.Details = artist[id-1]
	detail.Rel = relations.Index[id-1]
	data.Title = "relations"
	data.Data = detail

	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
