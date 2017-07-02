package main

import (
  "sync"
  "github.com/fatih/color"
)

type worker struct {
  id int
  workQueue chan workRequest
}

var workerQueue chan chan workRequest

var done int = 0

var doneMutex sync.Mutex

func (w *worker) start(){
  go func(){
    for{
      // Be like, "Hey, I'm available again"
      workerQueue <- w.workQueue
      select{
        // Receive a work request.
      case work := <-w.workQueue:

        // Time consuming job here
        //time.Sleep(2000 * time.Millisecond)
        work.execute()

        doneMutex.Lock()
        done++
        green := color.New(color.FgGreen, color.Bold)
        green.Printf("Completed #%d: Worker is done (description = %s)\n", done, work)
        doneMutex.Unlock()
      }
    }
  }()
}


func startWorkers(){

  workerQueue = make(chan chan workRequest, *nWorkers)

  for i := 0; i<*nWorkers; i++ {
    worker := worker{id: i+1, workQueue: make(chan workRequest)}
    worker.start()
  }
}
