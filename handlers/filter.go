package handlers

import "net/http"

func Filter(w http.ResponseWriter, r *http.Request) {
	var data Info
	data.Title = "filter"
	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}