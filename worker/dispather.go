package worker

import "fmt"

type Dispatcher struct {
  work_pool   chan chan Job
  len         int
  handle_type  string
}

func NewDispatcher(max_worker int, handle_type string) *Dispatcher {
  worker_pool := make(chan chan Job, max_worker)
  return  &Dispatcher{ work_pool:  worker_pool, len:  max_worker, handle_type: handle_type }
}

func (d *Dispatcher) Run() {
  // make three worker for process
  // todo:  use the same work_pool, so can limit in 3 process at the same time
  // todo: this will use job_channel put in work_pool first
  fmt.Println("make ", d.len, " workers for process jobs")
  for i := 0; i < d.len; i++ {
    worker := NewWorker(d.work_pool, d.handle_type)
    worker.Start()
  }

  go d.dispatcher()
}


// get the job_channel from work_pool,
func (d *Dispatcher) dispatcher() {
  for {
    select {
    case job := <- Job_queue:
      go func(job Job) {
        // fixme: 读取 work_pool，  Start() 方法写进 work_pool， 这里是读取 work_pool, 如果 work_pool 信道没有数据， 这里会堵塞
        //job_channel := <- d.work_pool
        job_channel := <- d.work_pool
        // fixme: 这里如果写成 _ := <- d.work_pool 会产生错误， 信道读取必须要赋予变量
        //_ := <- d.work_pool

        job_channel <- job
      }(job)
    }
  }
}
