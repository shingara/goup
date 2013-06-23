package web

import (
  "encoding/json"
  "log"
  "fmt"
  "net/http"

  "github.com/gorilla/mux"

  "github.com/shingara/goup/models"
)

func Start() {
  r := mux.NewRouter()
  r.HandleFunc("/urls", UrlsHandler)
  http.Handle("/", r)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func UrlsHandler(w http.ResponseWriter, req *http.Request) {
  urls := models.AllUrl()
  jsons, err := json.Marshal(urls)
  if err != nil {
    fmt.Println("error:", err)
  }
  w.Write(jsons)
}
