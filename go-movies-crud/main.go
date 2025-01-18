package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(movies) == 0 {
		json.NewEncoder(w).Encode([]Movie{})
		return
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]

	for _, item := range movies {
		if item.ID == movieId {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Movie Id not found"})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaition/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	movieId := params["id"]

	for index, item := range movies {
		if item.ID == movieId {
			// removing existing movie from slice
			movies = append(movies[:index], movies[index+1:]...)

			var updatedMovie Movie
			if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
				http.Error(w, "Invalida request payload", http.StatusBadRequest)
				return
			}

			updatedMovie.ID = movieId
			movies = append(movies, updatedMovie)

			// return the new updated movie only
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	for index, item := range movies {
		if item.ID == movieId {
			movies = append(movies[:index], movies[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "movie successfully deleted"})
			return
		}
	}

	// no movie found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "movie not found"})

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1", Isbn: "1234536", Title: "Movie 1", Director: &Director{
			Firstname: "firstname",
			Lastname:  "lastname",
		},
	})

	movies = append(movies, Movie{
		ID: "2", Isbn: "7890123", Title: "Movie 2", Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	// fmt.Println("Initially Movies", movies)

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
