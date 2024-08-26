package handlers

import (
	"net/http"
	"strconv"

	groupie "groupie/data"
)

type LocData struct {
	Details groupie.Artist
	Loc     groupie.Indexx
}

func LocationtHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var details LocData
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

	location, err := groupie.FetchLocation()
	artist, err1 := groupie.FetchArtist()
	if err != nil || err1 != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
	details.Details = artist[id-1]
	details.Loc = location.Index[id-1]

	if id > len(location.Index) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	data.Title = "location"
	data.Data = details

	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
