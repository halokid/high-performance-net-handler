package worker

import "fmt"

type Request struct {
  job Job
}


// make the jobs, len is max_job
func NewRequest(max_job int) (*Request)  {
  job := Job{pay_load:  Payload(max_job)}
  return &Request{job:  job}
}


// set the job to job_queue
func (r *Request) Run() {
  for i := 1; i < int(r.job.pay_load): i++ {
    job := Job{ pay_load:  Payload(i) }
    fmt.Println("put ---", i, "--- job into job_queue, job_queue only get one job every time ")
    Job_queue <- job
  }
}
