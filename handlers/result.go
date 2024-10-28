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
	albumFrm := r.Form.Get("albumDateFrom")
	albumTo := r.Form.Get("albumDateTo")
	minMembers, err := strconv.Atoi(r.Form.Get("membersMin"))
	maxMembers, err := strconv.Atoi(r.Form.Get("membersMax"))
	location, err := strconv.Atoi(r.Form.Get("location"))

	if err != nil {
		ErrorPage(w, r, http.StatusBadRequest, "Bad request")
	}
	
	result = CreationDate(creationFrm, creationTo)
	result = AlbumYear(result, albumFrm, albumTo)
}

func CreationDate(creationFrm, creationTo int) (result []groupie.Artist) {
	for _ , artist := range Artiste {
		if artist.CreationDate >= creationFrm && artist.CreationDate <= creationTo {
			result = append(result, artist)
		}
	}
	return
}

func AlbumYear(res []groupie.Artist, AlbumFrm, AlbumTo string) (result []groupie.Artist) {
	if res == nil {
		res = Artiste
	}
	for _ , artist := range res {
		if artist.FirstAlbum >= AlbumFrm && artist.FirstAlbum <=AlbumTo {
			result = append(result, artist)
		}
	}
	return
}

func Members(res []groupie.Artist, minMembers, maxMembers int) (result []groupie.Artist) {
	if res == nil {
		res = Artiste
	}
	for _ , artist := range res {
		if len(artist.Members) <= minMembers && len(artist.Members)>= maxMembers {
			result = append(result, artist)
		}
	}
	return
}
