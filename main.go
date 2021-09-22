package main

import (
	"net/http"
	//"github.com/gorilla/mux"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
}
