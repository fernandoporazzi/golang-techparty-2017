package main

import (
	"log"
	"net/http"

	"github.com/golang-techparty-2017/movie"
	"github.com/golang-techparty-2017/page"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/").
		HandlerFunc(page.Index)

	router.
		Methods("GET").
		Path("/movie").
		HandlerFunc(movie.GetAllMovies)

	router.
		Methods("POST").
		Path("/movie").
		HandlerFunc(movie.CreateMovie)

	log.Fatal(http.ListenAndServe(":3000", router))
}
