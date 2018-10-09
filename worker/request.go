package worker

import (
  "github.com/r00tjimmy/high-performance-net-handler/utils"
  //"fmt"
  "net/http"
  "github.com/r00tjimmy/high-performance-net-handler/handler"
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
  for i := 1; i < int(r.job.pay_load); i++ {
    job := Job{ pay_load:  Payload(i) }
    fmt.Println("put ---", i, "--- job into job_queue, job_queue only get one job every time ")
    Job_queue <- job
  }
}

**/



// HTTP listening
// TODO: get job from frontend
func (r *Request) Run() {
  //r.handle_type = "http"
  r.SetHandle()
}


func (r *Request) SetHandle() {
  if r.handle_type == "http" {
    http.HandleFunc("/hpnh", HttpHandle)
    http.HandleFunc("/hpnh_upload", HttpUpload)
    err := http.ListenAndServe(":8089", nil)
    utils.CheckErr(err)
  }
}


// decorator for handler
//var HandleDecorator =


// set response to  HTTP
// TODO: handler for web frontend
func HttpHandle(w http.ResponseWriter, r *http.Request) {
  //fmt.Println()
  utils.DebugLog("------------------------- [JOB START] -------------------------")
  utils.DebugLog("[JOB] --------- HTTP handle start -------- ")
  // if no error
  get_job := 1
  job := Job{pay_load: PayLoad(get_job) }
  utils.DebugLog("put --- 1 --- job into job_queue, job_queue only get one job every time ")
  Job_queue <- job

  //fmt.Fprintf(w, "handle http request")
  handler.HttpFrontSample(w, r)
  utils.DebugLog("------------------------- [JOB END] -------------------------")
}



// http upload frontend
func HttpUpload(w http.ResponseWriter, r *http.Request) {
  utils.DebugLog("[JOB START] HttpUpload -------------------------")
  //handler.HttpUploadHandle(w, r)
  handler.HttpUploadDateTimeHandle(w, r)
}





