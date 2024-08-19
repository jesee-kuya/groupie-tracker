package groupie

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// Struct for locations
type Indexx struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Location struct {
	Index []Indexx `json:"index"`
}

type Indexxx struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Dates struct {
	Index []Indexxx `json:"index"`
}

type Indexxxx struct {
	Id int `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []Indexxxx `json:"index"`
}