package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Movies -Storing our movies
var movies []Movie

func main() {
	port := ":8000"
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "3452134",
		Title: "The start",
		Director: &Director{
			FirstName: "David",
			LastName:  "Ochieng",
		},
	})

	r.HandleFunc("/movies", getMovies)
	r.HandleFunc("/movies/id", getMovie)
	r.HandleFunc("/movies", createMovies)
	r.HandleFunc("/movies/id", updateMovies)
	r.HandleFunc("/movies/id", deleteMovies)

	fmt.Printf("Starting server on port %v\n", port)

	log.Fatal(http.ListenAndServe(port, r))

}
