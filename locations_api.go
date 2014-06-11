package main

import (
  "net/http"
  "time"

  "github.com/go-martini/martini"
  "github.com/martini-contrib/encoder"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

func getLocation(params martini.Params, enc encoder.Encoder, db *mgo.Database) (int, []byte) {
  var location Location

  if !bson.IsObjectIdHex(params["id"]) {
    return http.StatusNotFound, []byte("Location with id \"" + params["id"] + "\" not found")
  }

  c := db.C("locations")
  err := c.FindId(bson.ObjectIdHex(params["id"])).One(&location)
  if err != nil {
    return http.StatusNotFound, []byte("Location with id \"" + params["id"] + "\" not found")
  }

  return http.StatusOK, encoder.Must(enc.Encode(location))
}

func deleteLocation(params martini.Params, enc encoder.Encoder, db *mgo.Database) (int, []byte) {
  var location Location

  if !bson.IsObjectIdHex(params["id"]) {
    return http.StatusNotFound, []byte("Location with id \"" + params["id"] + "\" not found")
  }

  c := db.C("locations")
  err := c.FindId(bson.ObjectIdHex(params["id"])).One(&location)
  if err != nil {
    return http.StatusNotFound, []byte("Location with id \"" + params["id"] + "\" not found")
  }

  err = c.RemoveId(bson.ObjectIdHex(params["id"]))

  return http.StatusOK, encoder.Must(enc.Encode(location))
}

func listLocations(enc encoder.Encoder, r *http.Request, db *mgo.Database) (int, []byte) {
  var locations []Location

  c := db.C("locations")
  err := c.Find(nil).All(&locations)
  if err != nil {
    return http.StatusInternalServerError, []byte("Impossible to retrieve the locations: " + err.Error())
  }

  return http.StatusOK, encoder.Must(enc.Encode(locations))
}

func createLocation(newLocation NewLocationForm, enc encoder.Encoder, db *mgo.Database) (int, []byte) {
  // transform the form struct into a real Location
  location := Location{
    ID: bson.NewObjectId(),
    Coordinates: Coordinates{
      Lat:  newLocation.Lat,
      Long: newLocation.Long,
    },
    Name:        newLocation.Name,
    Description: newLocation.Description,
    Visited:     newLocation.Visited,
    CreatedAt:   time.Now(),
  }

  // and try to insert it
  c := db.C("locations")
  err := c.Insert(&location)

  if err != nil {
    return http.StatusInternalServerError, []byte("Impossible to insert the location: " + err.Error())
  }

  return http.StatusCreated, encoder.Must(enc.Encode(location))
}
