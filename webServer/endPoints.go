package groupie

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	artists := fetchIndex()
	err = tpl.Execute(w, artists)
	if err != nil {
		log.Fatal(err)
	}
}

func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("Id")
	data := fetchArtist(id)

	tmpl, errtpl := template.ParseFiles("templates/artist.html")
	if errtpl != nil {
		log.Fatal(errtpl)
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
