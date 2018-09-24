package worker

import (
  "../utils"
  "fmt"
  "net/http"
)

type Request struct {
  job           Job
  handle_type   string
}


// make the jobs, len is max_job
func NewRequest(max_job int, handle_type string) (*Request)  {
  job := Job{pay_load:  PayLoad(max_job)}
  return &Request{job:  job, handle_type: handle_type}
}


/**
// set the job to job_queue, just for test request
func (r *Request) Run() {
  // fixme: 循环添加进去所有的job, pay_load 只是记录job的次序的， 表示正在处理到那个job， 并没实际意义
  for i := 1; i < int(r.job.pay_load); i++ {
    job := Job{ pay_load:  Payload(i) }
    fmt.Println("put ---", i, "--- job into job_queue, job_queue only get one job every time ")
    Job_queue <- job
  }
}

**/



// HTTP listening
func (r *Request) Run() {
  r.handle_type = "http"
  r.SetHandle()
}


func (r *Request) SetHandle() {
  if r.handle_type == "http" {
    http.HandleFunc("/hpnh", HttpHandle)
    err := http.ListenAndServe(":8089", nil)
    utils.CheckErr(err)
  }
}


func HttpHandle(w http.ResponseWriter, r *http.Request) {
  fmt.Println("HTTP handle start -------- ")
  // if no error
  job := Job{ PayLoad(1) }
  fmt.Println("put ---", 1, "--- job into job_queue, job_queue only get one job every time ")
  Job_queue <- job

  fmt.Fprintf(w, "handle http request")
}






