package movie

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
)

var (
	database *mgo.Database
)

func init() {
	session, err := mgo.Dial("mongodb://localhost:27017/techparty-golang")
	if err != nil {
		panic(err)
	}

	database = session.DB("techparty-golang")
}

// GetAllMovies retrieves all movies in the database
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	c := database.C("movie")
	result := Movies{}

	err := c.Find(nil).All(&result)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

// CreateMovie exported
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	// c := database.C("movie")

	var m Movie

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
