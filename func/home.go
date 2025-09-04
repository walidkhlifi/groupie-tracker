package groupietracker

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"html/template"
	"net/http"
)

type General struct {
 Liji Location 
 Laja Artist
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
nb:=r.URL.Query().Get("Id")

		data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/"+nb)
	if err != nil {
		http.Error(w, "rceive data", http.StatusInternalServerError)
		return
	}
	defer data.Body.Close()

	var ccc Artist

	er := json.NewDecoder(data.Body).Decode(&ccc)
	if er != nil {
		fmt.Println("errrrrrrrr")
		return
	}

	data1, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/"+nb)
	if err != nil {
		http.Error(w, "rceive data", http.StatusInternalServerError)
		return
	}
	defer data1.Body.Close()

	
	var genr Location

	er = json.NewDecoder(data1.Body).Decode(&genr)
	if er != nil {
		fmt.Println(er)
		return
	}
		dat := General{
		Liji: genr,
		Laja: ccc,
	}


fmt.Println(dat)
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
	

	if r.Method!="GET"{
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
		return
	}
    if r.URL.Path!="/"{
		http.Error(w,"pâge not found",http.StatusNotFound)
		return
	}




	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "rceive data", http.StatusInternalServerError)
		return
	}
	defer data.Body.Close()
	var HOME []Artist

	er := json.NewDecoder(data.Body).Decode(&HOME)
	if er != nil {
		fmt.Println("errrrrrrrr")
		return
	}
	// for _, artist := range artiss {
	// 	fmt.Printf("Artist Name: %s\n", artist.Name)
	// 	fmt.Printf("Members: %v\n", artist.Members) // طبع الـmembers
	// }
	tmpl, _ := template.ParseFiles("web/home.html")
	tmpl.Execute(w, HOME)
}


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