package groupie

import (
	"encoding/json"
	"net/http"
)

func Response(str string) (*http.Response, error) {
	url := "https://groupietrackers.herokuapp.com/api/" + str

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func FetchArtist() ([]Artist, error) {
	var artist []Artist
	response, err := Response("artists")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decode := json.NewDecoder(response.Body)
	jsErr := decode.Decode(&artist)
	if jsErr != nil {
		return nil, jsErr
	}
	return artist, nil
}

func FetchLocation() (Location, error) {
	var locations Location
	response, err := Response("locations")
	if err != nil {
		return locations, err
	}
	defer response.Body.Close()

	decode := json.NewDecoder(response.Body)
	jsErr := decode.Decode(&locations)
	if jsErr != nil {
		return locations, jsErr
	}
	return locations, nil
}

func FetchDates() (Dates, error) {
	var dates Dates
	response, err := Response("dates")
	if err != nil {
		return dates, err
	}
	defer response.Body.Close()

	decode := json.NewDecoder(response.Body)
	jsErr := decode.Decode(&dates)
	if jsErr != nil {
		return dates, jsErr
	}
	return dates, nil
}

func FetchRelation() (Relations, error) {
	var relation Relations
	response, err := Response("relation")
	if err != nil {
		return relation, err
	}
	defer response.Body.Close()

	decode := json.NewDecoder(response.Body)
	jsErr := decode.Decode(&relation)
	if jsErr != nil {
		return relation, jsErr
	}
	return relation, nil
}
