package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t := template.New("Label de ma template")
	t = template.Must(t.ParseFiles("View/index.html"))
	err := t.ExecuteTemplate(w, "index.html", ArtistsFull)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

func ArtistsPage(w http.ResponseWriter, r *http.Request) {
	// Recup values of the form
	SearchName := r.FormValue("search")
	StartDateArtists := r.FormValue("StartDateArtist")
	EndDateArtists := r.FormValue("StartEndArtist")
	StartDateAlbum := r.FormValue("StartDateAlbum")
	EndDateAlbum := r.FormValue("StartEndAlbum")
	Country := r.FormValue("country")
	Groups := r.FormValue("Groupes")
	Artistes := r.FormValue("Artistes")
	NewArtists := FilterArtistsByCreationDateOfAlbum(StartDateAlbum, EndDateAlbum, ArtistsFull)
	NewArtists = FilterArtistsByCreationDate(StartDateArtists, EndDateArtists, NewArtists)
	NewArtists = FilterArtistsByName(SearchName, NewArtists)
	NewArtists = FilterArtistByCountry(Country, NewArtists)
	NewArtists = DisplayArtistsByTypesWithForm(Groups, Artistes, NewArtists)
	t := template.New("Label de ma template")
	t = template.Must(t.ParseFiles("View/artists.html"))
	err := t.ExecuteTemplate(w, "artists.html", NewArtists)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}

}

func ConcertPage(w http.ResponseWriter, r *http.Request) {
	pattern := r.FormValue("search")
	NewArtists := ConcertFilterSearchAll(pattern, ArtistsFull)
	t := template.New("Label de ma template")
	t = template.Must(t.ParseFiles("View/concert.html"))
	err := t.ExecuteTemplate(w, "concert.html", NewArtists)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

func ArtistDetail(w http.ResponseWriter, r *http.Request) {
	CurrArtistID := r.FormValue("CurrentArtist")
	CurrArtist := ConvertToInt(CurrArtistID) - 1
	t := template.New("Label de ma template")
	t = template.Must(t.ParseFiles("View/ArtistDetail.html"))
	err := t.ExecuteTemplate(w, "ArtistDetail.html", ArtistsFull[CurrArtist])

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

func Billeterie(w http.ResponseWriter, r *http.Request) {
	t := template.New("Label de ma template")
	t = template.Must(t.ParseFiles("View/billeterie.html"))
	err := t.ExecuteTemplate(w, "billeterie.html", ArtistsFull)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}
