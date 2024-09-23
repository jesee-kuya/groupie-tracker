package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	groupie "groupie/data"
)

// Info holds the title and data to be rendered by the template.
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

// Handler routes incoming HTTP requests to the appropriate handler functions.
func Handler(w http.ResponseWriter, r *http.Request) {
	if Err5 != nil {
		Temp, Err5 = template.ParseGlob("template/*.html")
		return
	}
	if Err1 != nil || Err2 != nil || Err3 != nil || Err4 != nil {
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
	} else if r.URL.Path == "/search" {
		if r.Method == "GET" {
			SearchHandler(w, r)
		} else {
			ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
	} else {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
}

// ErrorPage generates an error page with the specified HTTP status code and message.
func ErrorPage(w http.ResponseWriter, r *http.Request, code int, msg string) {
	var data Info
	data.Title = "error"
	data.Data = fmt.Sprintf("Error: %v %v.", code, msg)
	w.WriteHeader(code)
	Temp.ExecuteTemplate(w, "base.html", data)
}

// GetId extracts and validates the artist ID from the query parameters.
func GetId(w http.ResponseWriter, r *http.Request) int {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil || id > len(Artiste) || id <= 0 {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return 0
	}
	return id
}
