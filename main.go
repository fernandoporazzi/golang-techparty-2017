package main

import (
	"log"
	"net/http"

	"github.com/golang-techparty-2017/movie"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/movie").
		HandlerFunc(movie.GetAllMovies)

	log.Fatal(http.ListenAndServe(":3000", router))
}
