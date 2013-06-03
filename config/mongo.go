package config

import (
  "labix.org/v2/mgo"
)

// Get a mongoDB session to this app
func Get_mongo_session() *mgo.Session{
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  /* defer session.Close() */
  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  return session
}

// Get a mongoDB Database to this app
func Get_mongo_db() *mgo.Database {
  return Get_mongo_session().DB("goup")
}
