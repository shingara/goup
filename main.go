package main

import (
  "fmt"
  "time"
  "os"

  "github.com/shingara/goup/config"
  "github.com/shingara/goup/request"
  "github.com/shingara/goup/models"
)

func main() {

  config.InitMongo()
  defer config.CloseSession()

  if (len(os.Args) > 1 && os.Args[1] == "add") {
    url_element := &url.Url{os.Args[2]}
    url.Add(url_element)
  } else {

    /* config.log_level(config.Level_info) */
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
        fmt.Printf("t : %v\n", t)
        for _, url := range url.All() {
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
}
