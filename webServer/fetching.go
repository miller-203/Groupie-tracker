package groupie

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func fetchIndex() []Artist {
	var artistians []Artist
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err1 := json.NewDecoder(resp.Body).Decode(&artistians)
	if err1 != nil {
		log.Fatal(err)
	}
	return artistians
}

func FetchData(ss interface{}, v string, s string, wg *sync.WaitGroup) {
	defer wg.Done()
	artist, errArt := http.Get("https://groupietrackers.herokuapp.com/api/" + v + s)
	if errArt != nil {
		log.Fatal(errArt)
	}
	defer artist.Body.Close()
	json.NewDecoder(artist.Body).Decode(&ss)
}

func fetchArtist(s string) Result {
	var art Artist
	var loc Locations
	var dat Dates
	var rel Relation
	var wg sync.WaitGroup

	wg.Add(4)
	go FetchData(&art, "artists/", s, &wg)
	go FetchData(&loc, "locations/", s, &wg)
	go FetchData(&dat, "dates/", s, &wg)
	go FetchData(&rel, "relation/", s, &wg)
	wg.Wait()
	result := Result{art, loc, dat, rel}

	return result
}
