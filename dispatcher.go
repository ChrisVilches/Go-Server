package main

import (
  "net/http"
  "github.com/fatih/color"
  "encoding/json"
)

type jsonData struct {
    Code string
}


func dispatch(work workRequest){
    worker := <-workerQueue
    worker <- work
    color.Cyan("### Work request was dispatched")
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    color.Red("Bad request: Not using POST")
    w.Header().Set("Allow", "POST")
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  decoder := json.NewDecoder(r.Body)
  var t jsonData
  err := decoder.Decode(&t)
  if err != nil {
    panic(err)
  }
  defer r.Body.Close()

  code := t.Code

  if len(code) < 1 {
    color.Red("Bad request: No code attribute")
    http.Error(w, "The request must have a code attribute.", http.StatusBadRequest)
    return
  }

  work := workRequest{code: code}
  go dispatch(work)

  green := color.New(color.FgGreen)
  green.Printf("### Work request was collected (code = %s)\n", code)

  w.WriteHeader(http.StatusCreated)
  return
}
