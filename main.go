package main

import (
	"fmt"
	"net/http"
	"groupietracker/func"
)

func main(){

http.HandleFunc("/", groupietracker.Home) 
http.HandleFunc("/artist", groupietracker.Artistt) 

fmt.Println("Server running at http://localhost:8080")
	ERR:=http.ListenAndServe(":8080",nil)
	if ERR!=nil{
		fmt.Println("err internel server")
	}
}