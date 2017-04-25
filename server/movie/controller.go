package movie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"crypto/md5"
	"encoding/hex"

	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
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

func getHash() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	hashInBytes := h.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes)
}

func uploadImage(h string, r *http.Request) {
	// 33.554432 MB
	r.ParseMultipartForm(32 << 20)

	file, _, err := r.FormFile("cover")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile("./uploads/"+h+".jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	if _, err := io.Copy(f, file); err != nil {
		panic(err)
	}
}

func insertMovie(h string, r *http.Request) Movie {
	c := database.C("movie")

	var m Movie

	y, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		panic(err)
	}

	m.ID = bson.NewObjectId()
	m.Name = r.FormValue("name")
	m.Year = y
	m.Synopsis = r.FormValue("synopsis")
	m.Cover = h + ".jpg"
	m.Categories = r.Form["category"]

	err = c.Insert(m)
	if err != nil {
		panic(err)
	}

	return m
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
	h := getHash()

	uploadImage(h, r)

	n := insertMovie(h, r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(n); err != nil {
		panic(err)
	}
}
