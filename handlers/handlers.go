package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	groupie "groupie/data"
)

var Temp, err = template.ParseGlob("template/*.html")

type Info struct {
	Title string
	Data  interface{}
}

func Handler(w http.ResponseWriter, r *http.Request) {
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
	} else {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	artist := groupie.FetchArtist()
	data.Title = "home"
	data.Data = artist

	Temp.ExecuteTemplate(w, "base.html", data)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}

	artist := groupie.FetchArtist()
	if id > len(artist) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	data.Title = "Artist"
	data.Data = artist[id-1]

	Temp.ExecuteTemplate(w, "base.html", data)
}

func ErrorPage(w http.ResponseWriter, r *http.Request, code int, msg string) {
	var data Info
	data.Title = "error"
	data.Data = fmt.Sprintf("Error: %v %v.", code, msg)
	w.WriteHeader(code)
	Temp.ExecuteTemplate(w, "base.html", data)
}

func LocationtHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}

	location := groupie.FetchLocation()
	if id > len(location.Index) {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	data.Title = "location"
	data.Data = location.Index[id-1]

	Temp.ExecuteTemplate(w, "base.html", data)
}
