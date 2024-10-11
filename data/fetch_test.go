package groupie

import (
	"net/http"
	"reflect"
	"testing"
)

func TestResponse(t *testing.T) {
	tests := []struct {
		name       string
		apiPath    string
		wantStatus int
		wantErr    bool
	}{
		// Test case for a valid API request
		{"Artists", "artists", http.StatusOK, false},
		{"Locations", "locations", http.StatusOK, false},
		{"Dates", "dates", http.StatusOK, false},
		{"Relation", "relation", http.StatusOK, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the Response function with the API path
			got, err := Response(tt.apiPath)

			// Check if we expect an error or not
			if (err != nil) != tt.wantErr {
				t.Errorf("Response() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If no error, check the HTTP status code of the response
			if got.StatusCode != tt.wantStatus {
				t.Errorf("Response() got status = %v, want %v", got.StatusCode, tt.wantStatus)
			}
		})
	}
}

func TestFetchArtist(t *testing.T) {

	type Artist struct {
		Id           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		ConcertDates string   `json:"concertDates"`
		Relations    string   `json:"relations"`
	}
	
	// var Artists struct{}
	tests := []struct {
		name    string
		want    Artist
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Rihanna", Artist{
			Id: 12, 
			Image: "https://groupietrackers.herokuapp.com/api/images/rihanna.jpeg", 
			Name: "Rihanna", 
			CreationDate: 2003, 
			Members: []string{0: "Robyn Rihanna Fenty"}, 
			FirstAlbum: "10-09-2005", 
			ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/12", 
			Relations: "https://groupietrackers.herokuapp.com/api/relation/12"}, 
		false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchArtist()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchArtist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var foundArtist Artist
			for _, gott := range got {
				if gott.Name == tt.want.Name {
					foundArtist = Artist(gott)
				}
			}
			if !reflect.DeepEqual(foundArtist, tt.want) {
				t.Errorf("FetchArtist() = %v\n, want %v", foundArtist, tt.want)
			}
		})
	}
}
