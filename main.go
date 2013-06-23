package main

import (
  "os"
  "github.com/shingara/goup/config"
  "github.com/shingara/goup/worker"
  "github.com/shingara/goup/request"
)

func main() {

  defer config.CloseSession()

  if (len(os.Args) > 1 && os.Args[1] == "add") {
    worker.AddUrl(os.Args[2])
  } else {
    request.Start()
  }
}
