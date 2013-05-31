package request

import(
  "fmt"
  "net/http"
  "log"
)

func Req(urls chan string, status chan int) {
  for url := range urls {
    fmt.Printf("nb urls %d\n", len(urls))
    fmt.Printf("launch request on %s\n", url)
    resp, err := http.Get(url)
    defer resp.Body.Close()

    if err != nil {
      log.Fatal(err)
    }
    status <- resp.StatusCode
  }
  close(status)
}

