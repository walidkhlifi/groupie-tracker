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
 Ddd Date
 rel Relation
}	


type Relation struct {
    Id      int      `json:"id"`
    Relations []string `json:"datesLocations"`
}

type Date struct{
	Id int  `json:"id"`
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
nb:=r.URL.Query().Get("Id")

if nb == "" {
    http.Error(w, "Missing or invalid ID", http.StatusBadRequest)
    return
}
//relation
data3, en := http.Get("https://groupietrackers.herokuapp.com/api/relation/"+nb)

	if en != nil {
		http.Error(w, "rceive data", http.StatusInternalServerError)
		return
	}
	defer data3.Body.Close()

	var rr Relation
	errr1 := json.NewDecoder(data3.Body).Decode(&rr)
	if errr1 != nil {
		fmt.Println("er4544")
		return
	}
fmt.Println(rr)

//date
data2, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/"+nb)

	if err != nil {
		http.Error(w, "rceive data", http.StatusInternalServerError)
		return
	}
	defer data2.Body.Close()

	var date Date

	errr := json.NewDecoder(data2.Body).Decode(&date)
	if errr != nil {
		fmt.Println("errrrrrrrr555")
		return
	}

	//artist
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

	// locations
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
		Ddd:date,
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
	

	if r.Method!="GET"{
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
		return
	}
    // if r.URL.Path!="/"{
	// 	http.Error(w,"pâge not found",http.StatusNotFound)
	// 	return
	// }




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