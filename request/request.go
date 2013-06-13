package request

import(
  "fmt"
  "net/http"
)


/* Wait url from the chan and return status in chan */
func Req(urls chan string, status chan int) {
  fmt.Printf("start worker\n")
  for url := range urls {
    fmt.Printf("launch request on %s\n", url)
    resp, err := http.Get(url)

    if err != nil {
      status <- 500
    } else {
      defer resp.Body.Close()
      status <- resp.StatusCode
    }
  }
  close(status)
  fmt.Printf("end worker\n")
}
