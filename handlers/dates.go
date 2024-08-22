package handlers

import (
	"net/http"
	"strconv"

	groupie "groupie/data"
)

func DatesHandler(w http.ResponseWriter, r *http.Request) {
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

	dates, err := groupie.FetchDates()
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
	if id > len(dates.Index) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	data.Title = "dates"
	data.Data = dates.Index[id-1]

	Temp.ExecuteTemplate(w, "base.html", data)
}
