package handlers

import (
	"net/http"

	groupie "groupie/data"
)

type LocData struct {
	Details groupie.Artist
	Loc     groupie.Indexx
}

func LocationtHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var details LocData

	id := GetId(w, r)
	if id <= 0 {
		return
	}

	details.Details = Artiste[id-1]
	details.Loc = Emplacement.Index[id-1]

	data.Title = "location"
	data.Data = details

	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}
