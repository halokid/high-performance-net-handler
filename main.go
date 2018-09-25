package main

import (
  //"./worker"
  "github.com/r00tjimmy/high-performance-net-handler/worker"
  "time"
)

var (
  max_worker = 3
  max_job = 10
  handle_type = "http"   // set network protocol type
)

func main() {
  // make the worker, listening work_pool channel
  dispatcher := worker.NewDispatcher(max_worker, handle_type)
  dispatcher.Run()

  // get requet
  request := worker.NewRequest(max_job, handle_type)
  request.Run()

  time.Sleep(time.Second * 10)
}

