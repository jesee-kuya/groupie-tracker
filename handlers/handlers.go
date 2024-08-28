package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	groupie "groupie/data"
)

type Info struct {
	Title string
	Data  interface{}
}

var (
	Artiste, Err1     = groupie.FetchArtist()
	Emplacement, Err2 = groupie.FetchLocation()
	Rapports, Err3    = groupie.FetchRelation()
	Tarehe, Err4      = groupie.FetchDates()
	Temp, Err5        = template.ParseGlob("template/*.html")
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if Err1 != nil || Err2 != nil || Err3 != nil || Err4 != nil || Err5 != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	if r.URL.Path == "/" || r.URL.Path == "/home" {
		if r.Method == "GET" {
			HomeHandler(w, r)
			return
		} else {
			ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
	} else if r.URL.Path == "/artist" {
		if r.Method == "GET" {
			ArtistHandler(w, r)
			return
		} else {
			ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
	} else if r.URL.Path == "/location" {
		if r.Method == "GET" {
			LocationtHandler(w, r)
			return
		} else {
			ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
	} else if r.URL.Path == "/dates" {
		if r.Method == "GET" {
			DatesHandler(w, r)
			return
		} else {
			ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
	} else if r.URL.Path == "/relations" {
		if r.Method == "GET" {
			RelationsHandler(w, r)
			return
		} else {
			ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
	} else {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
}

func ErrorPage(w http.ResponseWriter, r *http.Request, code int, msg string) {
	var data Info
	data.Title = "error"
	data.Data = fmt.Sprintf("Error: %v %v.", code, msg)
	w.WriteHeader(code)
	Temp.ExecuteTemplate(w, "base.html", data)
}

func GetId(w http.ResponseWriter, r *http.Request) int {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil || id > len(Artiste) || id < 0 {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return 0
	}
	return id
}
