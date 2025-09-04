package groupietracker

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"html/template"
	"net/http"
)

type General struct {
 jiji Location 
 jaja Artist
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
	data1, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/1")
	if err != nil {
		http.Error(w, "rceive data", http.StatusInternalServerError)
		return
	}
	defer data1.Body.Close()

	// 		data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	// 	if err != nil {
	// 		http.Error(w,"rceive data", http.StatusInternalServerError)
	// 		return
	// 	}
	// defer data.Body.Close()

	//		data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	//	if err != nil {
	//		http.Error(w,"rceive data", http.StatusInternalServerError)
	//		return
	//	}
	//
	// defer data.Body.Close()
	var genr General

	er := json.NewDecoder(data1.Body).Decode(&genr)
	if er != nil {
		fmt.Println(er)
		return
	}
	fmt.Println(genr)

	tmpl, _ := template.ParseFiles("web/artist.html")
	tmpl.Execute(w, genr)
}

// 2222
func Home(w http.ResponseWriter, r *http.Request) {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "rceive data", http.StatusInternalServerError)
		return
	}
	defer data.Body.Close()
	var artiss []Artist

	er := json.NewDecoder(data.Body).Decode(&artiss)
	if er != nil {
		fmt.Println("errrrrrrrr")
		return
	}
	// for _, artist := range artiss {
	// 	fmt.Printf("Artist Name: %s\n", artist.Name)
	// 	fmt.Printf("Members: %v\n", artist.Members) // طبع الـmembers
	// }
	tmpl, _ := template.ParseFiles("web/home.html")
	tmpl.Execute(w, artiss)
}
