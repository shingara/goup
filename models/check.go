package models

import (
  "github.com/shingara/goup/config"
  "labix.org/v2/mgo/bson"
  "time"
)

var (
  check_collection_name string
)

func init(){
  check_collection_name = "checks"
}

// A struct representing a Check
type Check struct {
  UrlId bson.ObjectId `bson:"url_id"`
  CheckDate time.Time
  Status int
}

// Insert this url with url pass in args
// TODO: Need valid Url
func AddCheck(check *Check)  {
  c := config.Config_collection(check_collection_name)
  defer c.Close()

  err := c.Collection.Insert(check)
  if err != nil {
    panic(err)
  }
}
