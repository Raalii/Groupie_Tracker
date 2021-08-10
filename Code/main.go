package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "./controllers"
)

func main() {
	controllers.GetAll()
	fs := http.FileServer(http.Dir("./View/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/artists", controllers.ArtistsPage)
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/concert", controllers.ConcertPage)
	http.HandleFunc("/artists/details", controllers.ArtistDetail)
	http.HandleFunc("/billeterie", controllers.Billeterie)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
