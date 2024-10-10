package handlers

import(
	"net/http"
	"strings"
	"encoding/json"
	groupie "groupie/data"
)

// SuggestionHandler handles suggestions for artists based on a query parameter.
func SuggestionHandler(w http.ResponseWriter, r *http.Request) {
    var results []groupie.Artist
    word := r.URL.Query().Get("query")

    for _, artist := range Artiste {
        if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(word)) {
            results = append(results, artist)
        }
    }

    // Return the results as JSON.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}
