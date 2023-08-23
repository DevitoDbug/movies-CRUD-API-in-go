package main

import (
	"encoding/json"
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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
}

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

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "9678963",
		Title: "The second",
		Director: &Director{
			FirstName: "Gerald",
			LastName:  "Bahati",
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
