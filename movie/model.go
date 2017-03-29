package movie

import (
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Year        int           `json:"year"`
}

type Movies []Movie
