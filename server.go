package main

import (
  "fmt"
  "net/http"

  "labix.org/v2/mgo"
)

func main() {
  // connect to the DB
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  session.SetMode(mgo.Monotonic, true)
  defer session.Close()

  // create the API server
  fmt.Println("Starting server on port 3000")
  http.ListenAndServe(":3000", ApiHandler(session))
}
