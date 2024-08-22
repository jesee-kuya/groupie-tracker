package handlers

import (
	"fmt"
	"net/http"
	"text/template"
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
