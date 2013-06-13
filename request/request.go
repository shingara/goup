package request

import(
  "fmt"
  "net/http"
  "time"
  "github.com/shingara/goup/models/url"
  "github.com/shingara/goup/models/check"
)

/* Wait url from the chan and return status in chan */
func Req(urls chan url.Url, status chan int) {
  fmt.Printf("start worker\n")
  for url := range urls {
    fmt.Printf("launch request on %s\n", url)
    resp, err := http.Get(url.Name)
    var status_code int = 500

    if err == nil {
      status_code = resp.StatusCode
      defer resp.Body.Close()
    }
    check_doc := &check.Check{
      UrlId: url.Id_,
      CheckDate: time.Now(),
      Status: status_code,
    }
    check.Add(check_doc)
    status <- check_doc.Status
  }
  close(status)
  fmt.Printf("end worker\n")
}
