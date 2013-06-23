package request

import(
  "time"
  "fmt"
  "github.com/shingara/goup/models"
)

func url_sender(urls chan models.Url) {
  ticker := time.NewTicker(time.Millisecond * 1000)
  for t := range ticker.C {
    fmt.Printf("t : %v\n", t)
    for _, url := range models.AllUrl() {
      urls <- url
    }
  }
}

func read_status(status chan int) {
  for statut := range status {
    fmt.Printf("status : %d\n", statut)
  }
}

func Start() {
  /* config.log_level(config.Level_info) */
  urls := make(chan models.Url, 100)
  status := make(chan int, 100)

  /* Launch request pool */
  for w := 1; w <= 10; w++ {
    go Req(urls, status)
  }

  /* Send urls */
  go url_sender(urls)
  /* Read status */
  go read_status(status)

  time.Sleep(time.Millisecond * 20000)
  close(urls)
}
