package main

import (
	"encoding/json"
	"fmt"
	"github.com/aguazul-marco/goProjects/go-movies-crud/data"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, movie := range data.Movies {
		if movie.ID == params["id"] {
			data.Movies = append(data.Movies[:i], data.Movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data.Movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range data.Movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie data.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	data.Movies = append(data.Movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, m := range data.Movies {
		if m.ID == params["id"] {
			data.Movies = append(data.Movies[:i], data.Movies[i+1:]...)
			var movie data.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			data.Movies = append(data.Movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	data.Movies = append(data.Movies, data.Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &data.Director{FirstName: "John", LastName: "Smith"}})
	data.Movies = append(data.Movies, data.Movie{ID: "2", Isbn: "438255", Title: "Movie Two", Director: &data.Director{FirstName: "Tim", LastName: "Jones"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
