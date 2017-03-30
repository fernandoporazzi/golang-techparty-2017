package movie

import (
	"gopkg.in/mgo.v2/bson"
)

// Movie is exported in order to create or update database documents
type Movie struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `json:"name"`
	Year       int           `json:"year"`
	Synopsis   string        `json:"synopsis"`
	Cover      string        `json:"cover"`
	Categories []string      `json:"categories"`
}

// Movies exported as collection of Movie
type Movies []Movie
