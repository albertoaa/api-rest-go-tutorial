package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contact", Contact)
	router.HandleFunc("/peliculas", MoviesList)
	router.HandleFunc("/pelicula/{id}", MovieShow)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la pagina de contacto")
}

func MoviesList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listado de Peliculas")
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}
