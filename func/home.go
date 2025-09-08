package groupietracker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type General struct {
	Liji Location
	Laja Artist
	Ddd  Date
	Rel  Relation
}

type Relation struct {
	Id        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

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

// /111
func Artistt(w http.ResponseWriter, r *http.Request) {
	nb := r.URL.Query().Get("Id")

	if nb == "" {
		http.Error(w, "Missing or invalid 'Id' parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(nb)
	if err != nil || id < 1 || id > 52 {
		http.Error(w, "Invalid 'Id' parameter. It must be an integer between 1 and 52.", http.StatusBadRequest)
		return
	}

	// relation
	data3, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + nb)
	if err != nil {
		http.Error(w, "failed to receive data", http.StatusInternalServerError)
		return
	}
	defer data3.Body.Close()

	var hp Relation
	err = json.NewDecoder(data3.Body).Decode(&hp)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		http.Error(w, "failed to decode data", http.StatusInternalServerError)
		return
	}

	// date
	data2, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + nb)
	if err != nil {
		http.Error(w, "failed to receive date data", http.StatusInternalServerError)
		return
	}
	defer data2.Body.Close()

	var date Date
	err = json.NewDecoder(data2.Body).Decode(&date)
	if err != nil {
		fmt.Println("error decoding date data:", err)
		http.Error(w, "failed to decode date data", http.StatusInternalServerError)
		return
	}

	// artist
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + nb)
	if err != nil {
		http.Error(w, "failed to receive artist data", http.StatusInternalServerError)
		return
	}
	defer data.Body.Close()

	var ccc Artist
	err = json.NewDecoder(data.Body).Decode(&ccc)
	if err != nil {
		fmt.Println("error decoding artist data:", err)
		http.Error(w, "failed to decode artist data", http.StatusInternalServerError)
		return
	}

	// locations
	data1, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + nb)
	if err != nil {
		http.Error(w, "failed to receive location data", http.StatusInternalServerError)
		return
	}
	defer data1.Body.Close()

	var genr Location
	err = json.NewDecoder(data1.Body).Decode(&genr)
	if err != nil {
		fmt.Println("error decoding location data:", err)
		http.Error(w, "failed to decode location data", http.StatusInternalServerError)
		return
	}
	/////////////////
	dat := General{
		Liji: genr,
		Laja: ccc,
		Ddd:  date,
		Rel:  hp,
	}

	tmpl, err := template.ParseFiles("web/artist.html")
	if err != nil {
		fmt.Println("Template Error:", err)
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, dat)
	if err != nil {
		fmt.Println("Execute Error:", err)
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}
}

// 2222
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "pâge not found", http.StatusNotFound)
		return
	}

	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}
	defer data.Body.Close()

	var HOME []Artist
	er := json.NewDecoder(data.Body).Decode(&HOME)
	if er != nil {
		fmt.Println("Failed to decode data")
		http.Error(w, "Failed to decode data", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("web/home.html")
	if err != nil {
		fmt.Println("Template Error:", err)
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, HOME)
	if err != nil {
		fmt.Println("Template Execution Error:", err)
		http.Error(w, "Template execution failed", http.StatusInternalServerError)
		return
	}
}
