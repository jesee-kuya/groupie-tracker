package handlers

import (
	groupie "groupie/data"
	"net/http"
)

func Results(w http.ResponseWriter, r *http.Request) {
	var result []groupie.Artist
	if err := r.ParseForm(); err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	creationFrm := r.Form.Get("creationDateFrom")
	creationTo := r.Form.Get("creationDateTo")
	albumFromYear := r.Form.Get("albumDateFrom")
	albumtoYear := r.Form.Get("albumDateTo")
	minMembers := r.Form.Get("membersMin")
	maxMembers := r.Form.Get("membersMax")
	location := r.Form.Get("location")
}

func CreationDate(creationFrm, creationTo int) (result []groupie.Artist) {
	for _ , artist := range Artiste {
		if artist.CreationDate >= creationFrm && artist.CreationDate <= creationTo {
			result = append(result, artist)
		}
	}
	return
}
