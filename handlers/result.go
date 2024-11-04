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
	creationFrm, _ := strconv.Atoi(r.Form.Get("creationDateFrom"))
	creationTo, _ := strconv.Atoi(r.Form.Get("creationDateTo"))
	albumFrm := r.Form.Get("albumDateFrom")
	albumTo := r.Form.Get("albumDateTo")
	location := r.Form.Get("location")
	memb1 := r.Form.Get("members1")
	memb2 := r.Form.Get("members2")
	memb3 := r.Form.Get("members3")
	memb4 := r.Form.Get("members4")
	memb5 := r.Form.Get("members5")
	memb6 := r.Form.Get("members6")
	memb7 := r.Form.Get("members7")
	memb8 := r.Form.Get("members8")
	memb9 := r.Form.Get("members9")
	memb10Plus := r.Form.Get("members10plus")

	members := CheckMembers(memb1, memb2, memb3, memb4, memb5, memb6, memb7, memb8, memb9, memb10Plus)

	result = CreationDate(creationFrm, creationTo)
	result = AlbumYear(result, albumFrm, albumTo)
	result = Members(result, members)
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

func Members(res []groupie.Artist, members []int) (result []groupie.Artist) {
	if res == nil {
		res = Artiste
	}
	for _, artist := range res {
		for _, num := range members {
			if len(artist.Members) == num {
				result = append(result, artist)
				break
			}
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

func CheckMembers(str ...string) (members []int) {
	for _, v := range str {
		if v != "" {
			n, err := strconv.Atoi(v)
			if err == nil {
				members = append(members, n)
			}
		}
	}
	return
}
