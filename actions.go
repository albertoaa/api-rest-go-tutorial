package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"Sin limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorcesse"},
	Movie{"A todo gas", 2005, "Juan Antonio"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la pagina de contacto")
}

func MoviesList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)
}
