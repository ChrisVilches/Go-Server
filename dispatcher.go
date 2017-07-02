package main

import "net/http"
import "github.com/fatih/color"


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

  code := r.FormValue("code")

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
