package main

import (
 "time"
  "./worker"
)

var (
  max_worker = 3
  max_job = 10
)


func main() {
  // dispatcher the requests, make worker to process the request
  dispatcher := worker.NewDispatcher(max_worker)
  dispatcher.Run()

  // fixme: make the request, just for test
  request := worker.NewRequest(max_job)
  request.Run()

  time.Sleep(time.Second * 10)
}






