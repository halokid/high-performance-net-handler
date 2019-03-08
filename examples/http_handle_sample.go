package main

import (
  "github.com/r00tjimmy/high-performance-net-handler/worker"
)

var (
  maxWorker = 3
  maxJob = 10
  handleType = "http"   // set network protocol type
)

func main() {
  // make the worker at first
  // make the worker, listening work_pool channel
  dispatcher := worker.NewDispatcher(maxWorker, handleType)
  dispatcher.Run()

  // make the job for test sample
  // get requet
  request := worker.NewRequest(maxJob, handleType)
  request.Run()

}

