package main

import (
  "github.com/r00tjimmy/high-performance-net-handler/worker"
)

var (
  max_worker = 3
  max_job = 10
  handle_type = "http"   // set network protocol type
)

func main() {
  // make the worker at first
  // make the worker, listening work_pool channel
  dispatcher := worker.NewDispatcher(max_worker, handle_type)
  dispatcher.Run()

  // make the job for test sample
  // get requet
  request := worker.NewRequest(max_job, handle_type)
  request.Run()

}

