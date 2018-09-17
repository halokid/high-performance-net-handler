package worker

import "fmt"

type Request struct {
  job Job
}

func NewRequest(max_job int) (*Request) {
  job := Job{pay_load:  PayLoad(max_job)}
  return &Request{job:  job}
}



func (r *Request) Run() {
  for i := 1; i < int(r.job.pay_load); i++ {
    // todo:假如 最大值为50， 那么就是 50 个job， 一共有 50个 pay_load
    job := Job{ pay_load:  PayLoad(i) }
    fmt.Println("接收到request然后添加job, 这是 第 ", i, "个调度")
    Job_queue <- job
  }
}
