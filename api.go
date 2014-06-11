package main

import (
  "net/http"

  "github.com/K-Phoen/negotiate"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/encoder"
  "labix.org/v2/mgo"
)

func ApiHandler(dbSession *mgo.Session) http.Handler {
  m := martini.Classic()

  // install custom middlewares
  m.Use(formatNegotiation())
  m.Use(db(dbSession))

  // define our routes
  m.Group("/locations", func(r martini.Router) {
    r.Get("/", listLocations)
    r.Post("/", binding.Bind(NewLocationForm{}), createLocation)

    r.Get("/:id", getLocation)
    r.Delete("/:id", deleteLocation)
  })

  return m
}

func formatNegotiation() martini.Handler {
  negotiators := make(map[string]encoder.Encoder)
  //negotiators["application/xml"] = encoder.XmlEncoder{}
  negotiators["application/json"] = encoder.JsonEncoder{}

  return negotiate.NegotiateFormat(negotiators)
}

func db(session *mgo.Session) martini.Handler {
  // each request will have its own session
  return func(c martini.Context) {
      s := session.Clone()
      c.Map(s.DB("GoThere"))
      defer s.Close()
      c.Next()
  }
}
