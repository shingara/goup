package worker

import (
  "fmt"
  "github.com/shingara/goup/models"
  "labix.org/v2/mgo/bson"
)

func AddUrl(name string) {
  url_element := &models.Url{
    Id_: bson.NewObjectId(),
    Name: name,
  }
  models.AddUrl(url_element)
  fmt.Printf("add url : %d\n", name)
}
