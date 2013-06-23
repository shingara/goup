package main

import (
  "os"
  "github.com/shingara/goup/config"
  "github.com/shingara/goup/worker"
  "github.com/shingara/goup/request"
  "github.com/shingara/goup/web"
)

func main() {

  defer config.CloseSession()

  if (len(os.Args) > 1 && os.Args[1] == "add") {
    worker.AddUrl(os.Args[2])
  } else {
    finish := make(chan int, 2)
    go request.Start(finish)
    web.Start()
    finish <- 0
  }
}
