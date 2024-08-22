package handlers

import (
	"net/http"
	"strconv"

	groupie "groupie/data"
)

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
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

	relations, err := groupie.FetchRelation()
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
	if id > len(relations.Index) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	data.Title = "relations"
	data.Data = relations.Index[id-1]

	Temp.ExecuteTemplate(w, "base.html", data)
}
