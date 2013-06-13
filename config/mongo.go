package config

import (
  "labix.org/v2/mgo"
  "fmt"
)

var(
  DB *mgo.Database
)

// Get a mongoDB session to this app
func InitMongo(){
  fmt.Printf("init Mongo START\n")
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  /* defer session.Close() */
  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)
  DB = session.DB("goup")
  fmt.Printf("init Mongo DONE\n")
}

func CloseSession() {
  DB.Session.Close()
}
