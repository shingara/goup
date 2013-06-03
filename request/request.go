package request

import(
  "fmt"
  "net/http"
  "log"
)


/* Wait url from the chan and return status in chan */
func Req(urls chan string, status chan int) {
  for url := range urls {
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
