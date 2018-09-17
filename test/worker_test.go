package test

import (
  "../worker"
  "testing"
  //"fmt"
  "time"
)

func TestNewWorker(t *testing.T) {
  work_pool := make(chan chan worker.Job, 5)
  worker := worker.NewWorker(work_pool)
  worker.Start()
  time.Sleep(time.Second * 5)
}