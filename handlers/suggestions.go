package handlers

import (
    "encoding/json"
    "net/http"
    "strings"
)

type SuggestionResult struct {
    ArtistId int    `json:"artistId"`
    Name     string `json:"name"`
    MatchType string `json:"matchType"`
    MatchedItem string `json:"matchedItem"`
}

func SuggestionHandler(w http.ResponseWriter, r *http.Request) {
    var results []SuggestionResult
    word := r.URL.Query().Get("query")
    lowercaseWord := strings.ToLower(word)

    for _, artist := range Artiste {
        matches := matchArtist(artist, lowercaseWord)
        for _, match := range matches {
            results = append(results, SuggestionResult{
                ArtistId: artist.Id,
                Name: artist.Name,
                MatchType: match.MatchType,
                MatchedItem: match.MatchedItem,
            })
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}