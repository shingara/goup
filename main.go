package main

import (
  "fmt"
  "github.com/shingara/goup/request"
  "github.com/shingara/goup/models"
  "time"
)

func main() {
  urls := make(chan string, 100)
  status := make(chan int, 100)

  /* Launch request pool */
  for w := 1; w <= 10; w++ {
    go request.Req(urls, status)
  }

  ticker := time.NewTicker(time.Millisecond * 1000)

  /* Send urls */
  go func(){
    for t := range ticker.C {
      fmt.Printf("t : %d\n", t)
      for _, url := range model.Get_urls() {
        urls <- url.Name
      }
    }
  }()

  /* Read status */
  go func(){
    for statut := range status {
      fmt.Printf("status : %d\n", statut)
    }
  }()

  time.Sleep(time.Millisecond * 20000)
  close(urls)
}
