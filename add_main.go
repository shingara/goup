package main

import(
  "os"
  "github.com/shingara/goup/models"
)

func main() {
  url := &model.Url{os.Args[1]}
  model.Add_url(url)
}
