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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		data.Title = "error"
		data.Data = "500 Internal Server Error"
		Temp.ExecuteTemplate(w, "base.html", data)
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
		w.WriteHeader(http.StatusInternalServerError)
		data.Title = "error"
		data.Data = "500 Internal Server Error"
		Temp.ExecuteTemplate(w, "base.html", data)
		return
	}

	if r.Method != "GET" {
		fmt.Println(r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	artist := groupie.FetchArtist()
	data.Title = "Artist"
	data.Data = artist[id-1]

	Temp.ExecuteTemplate(w, "base.html", data)
}
