package model

import (
  "github.com/shingara/goup/config"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

// A struct representing an URL
type Url struct {
  Name string
}

// The collection where urls document are save
func Get_collection() *mgo.Collection {
  return config.Get_mongo_db().C("url")
}

// Get the url with this url pass in args
func Get_url(url string) Url {
  result := Url{}
  err := Get_collection().Find(bson.M{"url": url}).One(&result)
  if err != nil {
    panic(err)
  }
  return result
}

func Get_urls() []Url {
  result := []Url{}
  Get_collection().Find(nil).All(&result)
  return result
}

// Insert this url with url pass in args
// TODO: Need valid Url
func Add_url(url *Url)  {
  err := Get_collection().Insert(url)
  if err != nil {
    panic(err)
  }
}
