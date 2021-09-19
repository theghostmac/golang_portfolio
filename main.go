package main

import (
	"net/http"
	"log"
	"html/template"
	
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate, _ := template.ParseFiles("templates/home.html")
	homeTemplate.Execute(w, nil)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	fileServer := http.FileServer(http.Dir("static"))
	router.PathPrefix("/").Handler(http.StripPrefix("/static/", fileServer))
	log.Fatalln(http.ListenAndServe(":3000", nil))
}