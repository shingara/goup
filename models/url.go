package models

import (
  "github.com/shingara/goup/config"
  "labix.org/v2/mgo/bson"
)

// A struct representing an URL
type Url struct {
  Id_ bson.ObjectId `bson:"_id"`
  Name string
}

var (
  url_collection_name string
)

func init(){
  url_collection_name = "urls"
}

// Get the url with this url pass in args
func GetUrl(url string) *Url {
  c := config.Config_collection(url_collection_name)
  defer c.Close()

  result := &Url{}
  err := c.Collection.Find(bson.M{"url": url}).One(&result)
  if err != nil {
    panic(err)
  }
  return result
}

func AllUrl() []Url {
  c := config.Config_collection(url_collection_name)
  defer c.Close()

  result := []Url{}
  c.Collection.Find(nil).All(&result)
  return result
}

// Insert this url with url pass in args
// TODO: Need valid Url
func AddUrl(url *Url)  {
  c := config.Config_collection(url_collection_name)
  defer c.Close()

  err := c.Collection.Insert(url)
  if err != nil {
    panic(err)
  }
}
