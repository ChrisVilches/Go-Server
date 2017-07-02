package main

import (
  "flag"
  "fmt"
  "net/http"
)

var nWorkers = flag.Int("n", 4, "The number of workers to start")

func main() {
  flag.Parse()

  startWorkers()

  http.HandleFunc("/work", handler)

  if err := http.ListenAndServe(":8000", nil); err != nil {
    fmt.Println(err.Error())
  }
}
