package movie

import (
	// "fmt"
	"net/http"
	// "log"
	"encoding/json"
	"gopkg.in/mgo.v2"
)

var (
	Database *mgo.Database
)

func init()  {
	session, err := mgo.Dial("mongodb://localhost:27017/techparty-golang")
	if err != nil {
		panic(err)
	}

   	Database = session.DB("techparty-golang")
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	c := Database.C("movie")
	result := Movies{}

	err := c.Find(nil).All(&result)

	if err != nil {
		panic(err)
	}

	if err = json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
