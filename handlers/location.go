package handlers

import (
	"net/http"
	"strconv"

	groupie "groupie/data"
)

func LocationtHandler(w http.ResponseWriter, r *http.Request) {
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

	location := groupie.FetchLocation()
	if id > len(location.Index) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	data.Title = "location"
	data.Data = location.Index[id-1]

	Temp.ExecuteTemplate(w, "base.html", data)
}
