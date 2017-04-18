package main

import (
	"log"
	"net/http"

	"github.com/golang-techparty-2017/server/movie"
	"github.com/golang-techparty-2017/server/page"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// static handler
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads/"))))

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


  corsObj := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":8001", handlers.CORS(corsObj)(router)))
}
