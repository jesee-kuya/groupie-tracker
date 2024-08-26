package handlers

import (
	"net/http"
	"strconv"

	groupie "groupie/data"
)

type DateData struct {
	Details groupie.Artist
	Dt      groupie.Indexxx
}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var details DateData
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

	dates, err := groupie.FetchDates()
	artist, err1 := groupie.FetchArtist()

	if err != nil || err1 != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
	if id > len(dates.Index) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	details.Details = artist[id-1]
	details.Dt = dates.Index[id-1]
	data.Title = "dates"
	data.Data = details

	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
