package request_test

import(
  "testing"
  "github.com/stretchrcom/testify/assert"
  "github.com/shingara/goup/request"
)

func TestReq(t *testing.T) {
  urls := make(chan string, 3)
  status := make(chan int, 3)
  go request.Req(urls, status)
  urls <- "http://google.com"
  assert.Equal(t, <- status, 200)
}
