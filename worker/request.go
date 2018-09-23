package worker

import (
  "../handler"
  "../utils"
  "net/http"
)

type Request struct {
  job Job
}


// make the jobs, len is max_job
func NewRequest(max_job int) (*Request)  {
  job := Job{pay_load:  Payload(max_job)}
  return &Request{job:  job}
}


/**
// set the job to job_queue, just for test request
func (r *Request) Run() {
  for i := 1; i < int(r.job.pay_load); i++ {
    job := Job{ pay_load:  Payload(i) }
    fmt.Println("put ---", i, "--- job into job_queue, job_queue only get one job every time ")
    Job_queue <- job
  }
}

**/



// HTTP listening
func (r *Request) Run() {
  http.HandleFunc("/hpnh", handler.HttpHandle)
  err := http.ListenAndServe(":8089", nil)
  utils.CheckErr(err)

}




