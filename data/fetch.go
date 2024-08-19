package groupie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func FetchArtist() (artist []Artist) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	decode := json.NewDecoder(response.Body)
	jsErr := decode.Decode(&artist)
	if jsErr != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func FetchLocation() (locations Location) {
	url := "https://groupietrackers.herokuapp.com/api/locations"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	decode := json.NewDecoder(response.Body)
	jsErr := decode.Decode(&locations)
	if jsErr != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func FetchDates() (dates Dates) {
	url := "https://groupietrackers.herokuapp.com/api/dates"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	decode := json.NewDecoder(response.Body)
	jsErr := decode.Decode(&dates)
	if jsErr != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func FetchRelation() (relation Relations) {
	url := "https://groupietrackers.herokuapp.com/api/relation"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	decode := json.NewDecoder(response.Body)
	jsErr := decode.Decode(&relation)
	if jsErr != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
