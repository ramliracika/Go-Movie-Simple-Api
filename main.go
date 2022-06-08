package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

type Movies struct {
	Id       string     `json:"id"`
	Isbn     int64      `json:"isbn"`
	Title    string     `json:"title"`
	Year     int32      `json:"year"`
	Director *Directors `json:"director"`
}

type Directors struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(movies)
	}

}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movies
			_ = json.NewDecoder(r.Body).Decode(&movies)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies

	_ = json.NewDecoder(r.Body).Decode(&movies)
	movie.Id = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

var movies []Movies

func main() {
	movies = append(movies, Movies{Id: "1", Isbn: 34632423, Title: "Light Year", Year: 2022, Director: &Directors{Name: "Andy", Age: 31, Address: "California"}})
	movies = append(movies, Movies{Id: "2", Isbn: 234536423, Title: "Top Gun", Year: 2022, Director: &Directors{Name: "Smith", Age: 42, Address: "Silicon Valley"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", addMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Succes launch at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
