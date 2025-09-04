package groupietracker

import (
	// "fmt"
	"net/http"
)

func Interne(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func Notf(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

}
func Badreq(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)

}