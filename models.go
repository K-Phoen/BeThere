package main

import (
  "time"

  "labix.org/v2/mgo/bson"
)

type Coordinates struct {
  Lat  float64 `json:"latitude"`
  Long float64 `json:"longitude"`
}

type Location struct {
  ID bson.ObjectId `bson:"_id,omitempty" json:"id"`

  Name        string `json:"name"`
  Description string `json:"description"`

  Visited bool `json:"visited"`

  Coordinates Coordinates `json:"coordinates"`

  CreatedAt time.Time `json:"createdAt"`
}
