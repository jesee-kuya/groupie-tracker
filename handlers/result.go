package handlers

import (
	"net/http"
	"strconv"
	"strings"

	groupie "groupie/data"
)

func Results(w http.ResponseWriter, r *http.Request) {
	var result []groupie.Artist
	var data Info
	if err := r.ParseForm(); err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	creationFrm, err := strconv.Atoi(r.Form.Get("creationDateFrom"))
	creationTo, err1 := strconv.Atoi(r.Form.Get("creationDateTo"))
	albumFrm := r.Form.Get("albumDateFrom")
	albumTo := r.Form.Get("albumDateTo")
	minMembers, err2 := strconv.Atoi(r.Form.Get("membersMin"))
	maxMembers, err3 := strconv.Atoi(r.Form.Get("membersMax"))
	location := r.Form.Get("location")

	if err != nil || err1 != nil || err2 != nil || err3 != nil {
		ErrorPage(w, r, http.StatusBadRequest, "Bad request")
		return
	}

	result = CreationDate(creationFrm, creationTo)
	result = AlbumYear(result, albumFrm, albumTo)
	result = Members(result, minMembers, maxMembers)
	result = SearchLocation(result, location)

	data.Title = "result"
	data.Data = result

	if result == nil {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	Temp.ExecuteTemplate(w, "base.html", data)
}

func CreationDate(creationFrm, creationTo int) (result []groupie.Artist) {
	for _, artist := range Artiste {
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
	for _, artist := range res {
		if artist.FirstAlbum >= AlbumFrm && artist.FirstAlbum <= AlbumTo {
			result = append(result, artist)
		}
	}
	if result == nil {
		result = res
	}
	return
}

func Members(res []groupie.Artist, minMembers, maxMembers int) (result []groupie.Artist) {
	if res == nil {
		res = Artiste
	}
	for _, artist := range res {
		if len(artist.Members) <= minMembers && len(artist.Members) >= maxMembers {
			result = append(result, artist)
		}
	}
	if result == nil {
		result = res
	}
	return
}

func SearchLocation(res []groupie.Artist, location string) (result []groupie.Artist) {
	if res == nil {
		res = Artiste
	}
	if location == "" {
		return res
	}
	for _, artist := range res {
		for _, locale := range Emplacement.Index {
			if locale.Id == artist.Id {
				for _, loc := range locale.Locations {
					if strings.Contains(strings.ToLower(loc), location) {
						result = append(result, artist)
						break
					}
				}
				break
			}
		}
	}
	if result == nil {
		result = res
	}
	return
}
