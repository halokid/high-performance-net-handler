package worker

import (
  "fmt"
  "github.com/r00tjimmy/high-performance-net-handler/handler"
)

type Job struct {
  pay_load    PayLoad
}

type PayLoad int

/**
func (p PayLoad) Do() (err error) {
  fmt.Print("payload Do() process at -------", int(p), " job\n\n")
  err = nil
  return err
}
**/

func (p PayLoad) HttpDo() (err error) {
  handler.HttpProcess()
  err = nil
  return err
}

type Worker struct {
  work_pool       chan chan Job
  job_channel     chan Job
  quit            chan bool
  handle_type     string
}


// set the job queue for request, one job every time
var Job_queue chan Job
func init() {
  Job_queue = make(chan Job, 1)
}

func NewWorker(work_pool chan chan Job, handle_type string) Worker {
  return Worker{
    work_pool:      work_pool,
    job_channel:    make(chan Job),
    quit:           make(chan bool),
    handle_type:    handle_type,
  }
}


func (w *Worker) Start() {
  go func() {
    for {
      // use job_channel put in work_pool
      w.work_pool <- w.job_channel
      select {
      case job := <- w.job_channel:
        if w.handle_type == "http" {
          if err := job.pay_load.HttpDo(); err != nil {
            fmt.Println("[ERROR]---- payload Do() ", err.Error())
          }
        }

      case <- w.quit:
        return
      }
    }
  }()
}













