package handlers

import (
	"net/http"
	"strings"

	groupie "groupie/data"
)

// SearchHandler handles search requests for artists based on a query parameter.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var results []groupie.Artist
	word := r.URL.Query().Get("query")

	for _, artist := range Artiste {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(word)) {
			results = append(results, artist) // Add matching artist to results.
		}
	}

	if results == nil {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}
	data.Title = "search"
	data.Data = results
	Temp.ExecuteTemplate(w, "base.html", data)
}
