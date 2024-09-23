package handlers

import (
	"net/http"
)

// HomeHandler handles requests to the home page of the application.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	data.Title = "home"
	data.Data = Artiste
	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
