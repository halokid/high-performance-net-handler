package worker

import "fmt"

type Dispatcher struct {
  worker_pool     chan chan Job
  len             int
}


// make worker to process request
func NewDispatcher(max_worker int) *Dispatcher {
  worker_pool := make(chan chan Job, max_worker)
  return &Dispatcher{worker_pool:  worker_pool, len: max_worker}
}


func (d *Dispatcher) Run() {
  for i := 0; i < d.len; i++ {
    // fixme: 主要是分三个 worker 去处理进去 Job_queue 所有的job, 每条worker都调用 Start() 方法， 所以每条worker都是独立处理一个 channel， 每条channel处理完一个job之后， 会去抢Job_queue里面的job， 继续处理

    // fixme: 这里还有一个细节， d.work_pool 都是同一个 channel，所以可以控制到最多可以写入 3 个 job_channel 进去 这个 work_pool
    worker := NewWorker(d.worker_pool)
    worker.Start()
  }
  fmt.Println("[finished one work_pool process] len is  ---------", int(d.len), "\n\n")

  go d.dispatcher()
}


func (d *Dispatcher) dispatcher() {
  for {
    select {
    // block listening Job_queue
    case job := <- Job_queue:
      go func(job Job) {
        // block listening d.worker_pool

        // fixme： 然后是这里读取  d.work_pool, 实际上这里都是一个数据位置， 和 worker.Start那里的  w.work_pool, 靠 这个 一写一读， 实现了 work_pool 的处理管理，因为处理管理的大小是3， 所以每次最多只能生成 3 个 chan job, 所以最多只能异步处理 3 个request
        job_channel := <- d.worker_pool

        // from Job_queue read job, put it into job_channel
        job_channel <- job
      }(job)
    }
  }

}


