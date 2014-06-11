package main

import (
  "time"

  "labix.org/v2/mgo/bson"
)

type Coordinates struct {
  Lat  float64
  Long float64
}

type Location struct {
  ID bson.ObjectId `bson:"_id,omitempty"`

  Name        string
  Description string

  Visited bool

  Coordinates Coordinates

  CreatedAt time.Time
}
