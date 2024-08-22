package main

import (
	"log"
	"net/http"

	web "groupie/webServer"
)

func main() {
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/artist", web.ArtistInfo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
