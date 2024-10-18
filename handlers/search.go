package handlers

import (
	"net/http"
	"strconv"
	"strings"

	groupie "groupie/data"
)

type SearchResult struct {
	Artist      groupie.Artist
	MatchType   string
	MatchedItem string
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var data Info
	var results []groupie.Artist
	var newResults []groupie.Artist
	var ids []int
	word := r.URL.Query().Get("query")
	lowercaseWord := strings.ToLower(word)

	for _, artist := range Artiste {
		matches := matchArtist(artist, lowercaseWord)
		for _, match := range matches {
			results = append(results, match.Artist)
		}
	}
	if results == nil {
		ErrorPage(w, r, http.StatusNotFound, "Not found")
		return
	}

	for _, v := range results {
		if !Contains(v.Id, ids) {
			ids = append(ids, v.Id)
			newResults = append(newResults, v)
		}
	}
	data.Title = "search"
	data.Data = newResults
	Temp.ExecuteTemplate(w, "base.html", data)
}

func matchArtist(artist groupie.Artist, query string) []SearchResult {
	var matches []SearchResult
	seenMatches := make(map[string]bool) // Track unique matches

	// Check artist name
	if strings.Contains(strings.ToLower(artist.Name), query) {
		key := "Artist-" + artist.Name
		if !seenMatches[key] {
			matches = append(matches, SearchResult{Artist: artist, MatchType: "Artist", MatchedItem: artist.Name})
			seenMatches[key] = true
		}
	}

	// Check members
	for _, member := range artist.Members {
		if strings.Contains(strings.ToLower(member), query) {
			key := "Member-" + member
			if !seenMatches[key] {
				matches = append(matches, SearchResult{Artist: artist, MatchType: "Member", MatchedItem: member})
				seenMatches[key] = true
			}
		}
	}

	// Check dates
	if matchesDate(artist.FirstAlbum, query) {
		key := "FirstAlbum-" + artist.FirstAlbum
		if !seenMatches[key] {
			matches = append(matches, SearchResult{Artist: artist, MatchType: "First Album", MatchedItem: artist.FirstAlbum})
			seenMatches[key] = true
		}
	}

	if matchesDate(strconv.Itoa(artist.CreationDate), query) {
		key := "CreationDate-" + strconv.Itoa(artist.CreationDate)
		if !seenMatches[key] {
			matches = append(matches, SearchResult{Artist: artist, MatchType: "Creation Date", MatchedItem: strconv.Itoa(artist.CreationDate)})
			seenMatches[key] = true
		}
	}

	// Check locations
	for _, location := range Emplacement.Index {
		if location.Id == artist.Id {
			for _, loc := range location.Locations {
				if strings.Contains(strings.ToLower(loc), query) {
					key := "Location-" + loc
					if !seenMatches[key] {
						matches = append(matches, SearchResult{Artist: artist, MatchType: "Location", MatchedItem: loc})
						seenMatches[key] = true
					}
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

// Contains checks if an integer is present in a slice of integers
func Contains(id int, ids []int) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
