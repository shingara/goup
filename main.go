package main

import (
  "os"
  "fmt"
  "github.com/shingara/goup/request"
)

func main() {
  urls := make(chan string, 100)
  status := make(chan int, 100)

  go request.Req(urls, status)
  go request.Req(urls, status)
  go request.Req(urls, status)

  for w := 1; w <= 50; w++ {
    fmt.Println("send url to goroutine")
    urls <- os.Args[1]
  }
  fmt.Println("close urls")
  close(urls)
  for statut := range status {
    fmt.Printf("status : %d\n", statut)
  }
}
