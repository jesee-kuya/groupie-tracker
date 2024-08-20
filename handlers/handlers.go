package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	groupie "groupie/data"
)

var Temp, err = template.ParseGlob("template/*.html")

type Info struct {
	Title string
	Data  interface{}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	artist := groupie.FetchArtist()
	var data Info
	data.Title = "home"
	data.Data = artist

	Temp.Execute(w, data)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	artist := groupie.FetchArtist()
	var data Info
	data.Title = "Artist"
	data.Data = artist

	Temp.Execute(w, data)
}
