package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	// "github.com/golang-techparty-2017/controllers"
	"./movie"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/movie").
		HandlerFunc(movie.GetAllMovies)

	log.Fatal(http.ListenAndServe(":3000", router))
}
