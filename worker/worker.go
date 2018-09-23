package worker

type Payload int

type Job struct {
  pay_load      Payload
}

type Worker struct {
  work_pool      chan chan Job
  job_channel    chan Job
  quit           chan bool
}



// set the job queue for request
var Job_queue chan Job
func init() {
  Job_queue = make(chan Job, 1)
}


