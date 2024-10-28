package handlers

import (
	groupie "groupie/data"
	"net/http"
	"strconv"
)

func Results(w http.ResponseWriter, r *http.Request) {
	var result []groupie.Artist
	if err := r.ParseForm(); err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	creationFrm, err := strconv.Atoi(r.Form.Get("creationDateFrom"))
	creationTo, err := strconv.Atoi(r.Form.Get("creationDateTo"))
	albumFromYear, err:= strconv.Atoi(r.Form.Get("albumDateFrom"))
	albumtoYear, err := strconv.Atoi(r.Form.Get("albumDateTo"))
	minMembers, err := strconv.Atoi(r.Form.Get("membersMin"))
	maxMembers, err := strconv.Atoi(r.Form.Get("membersMax"))
	location, err := strconv.Atoi(r.Form.Get("location"))

	if err != nil {
		ErrorPage(w, r, http.StatusBadRequest, "Bad request")
	}
	
	result = CreationDate(creationFrm, creationTo)
}

func CreationDate(creationFrm, creationTo int) (result []groupie.Artist) {
	for _ , artist := range Artiste {
		if artist.CreationDate >= creationFrm && artist.CreationDate <= creationTo {
			result = append(result, artist)
		}
	}
	return
}

func AlbumYear(res []groupie.Artist, AlbumFrm, AlbumTo int) (result []groupie.Artist) {
	if res == nil {
		res = Artiste
	}
	for _ , artist := range res {
		if artist.CreationDate >= AlbumFrm && artist.CreationDate <=AlbumTo {
			result = append(result, artist)
		}
	}
	return
}
