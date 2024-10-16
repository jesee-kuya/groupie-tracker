package handlers

import (
    "net/http"
    "strings"
    "strconv"
    groupie "groupie/data"
)

type SearchResult struct {
    Artist groupie.Artist
    MatchType string
    MatchedItem string
}

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

func matchArtist(artist groupie.Artist, query string) []SearchResult {
    var matches []SearchResult

    if strings.Contains(strings.ToLower(artist.Name), query) {
        matches = append(matches, SearchResult{Artist: artist, MatchType: "Artist", MatchedItem: artist.Name})
    }

    for _, member := range artist.Members {
        if strings.Contains(strings.ToLower(member), query) {
            matches = append(matches, SearchResult{Artist: artist, MatchType: "Member", MatchedItem: member})
        }
    }

    if matchesDate(artist.FirstAlbum, query) {
        matches = append(matches, SearchResult{Artist: artist, MatchType: "First Album", MatchedItem: artist.FirstAlbum})
    }

    if matchesDate(strconv.Itoa(artist.CreationDate), query) {
        matches = append(matches, SearchResult{Artist: artist, MatchType: "Creation Date", MatchedItem: strconv.Itoa(artist.CreationDate)})
    }
    
    // Search through locations
    for _, location := range Emplacement.Index {
        if location.Id == artist.Id {
            for _, loc := range location.Locations {
                if strings.Contains(strings.ToLower(loc), query) {
                    matches = append(matches, SearchResult{Artist: artist, MatchType: "Location", MatchedItem: loc})
                }
            }
            break
        }
    }
    
    return matches
}

func matchesDate(date string, query string) bool {
    return strings.Contains(strings.ToLower(date), query)
}