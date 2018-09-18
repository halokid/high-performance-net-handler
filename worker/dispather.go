package worker

import (
  "fmt"
  //"time"
)

/**
负责 dispather 的逻辑代码
 */


type Dispatcher struct {
  worker_pool   chan chan Job
  len           int
}

func NewDispatcher(max_worker int) *Dispatcher {
  worker_pool := make(chan chan Job, max_worker)
  return  &Dispatcher{ worker_pool: worker_pool, len: max_worker }
}


/**
写入 worker
根据 dispather 生成 worker，每个worker 里面可能包含很多job, 这个要看 job_queue设置的长度
 */
func (d *Dispatcher) Run() {
  fmt.Println("work_pool 现在的长度是 ----------- ", len(d.worker_pool), ", 开始创建worker -----\n\n")
  for i := 0; i < d.len; i++ {
    fmt.Println("dispatcher 生成新的worker去处理job ------- ", i)
    fmt.Println("新建立的worker， ------- ", i)
    fmt.Println("建立的新的worker， work_pool的长度竟然也是max_worker的数值，这是不是有问题？----- ", i)

    worker := NewWorker(d.worker_pool)
    //worker := NewWorker(10)
    fmt.Println("新的worker生成完毕 -------- ", i)
    worker.Start()
    fmt.Println("调用第 ", i, "个worker的Start() 方法开始处理请求, 等待写入该worker的job" )
    fmt.Println("\n\n\n\n")

  }

  fmt.Println("执行到这里 xxxxx -----------------------------")
  // 把 job 写入worker 才算真正的开始执行处理请求
  go d.dispatcher()
  fmt.Println("执行到这里 yyyy -----------------------------")

  //time.Sleep(time.Second * 5)
}


/**
上面是生成worker的， 这里是写入 worker的 job的
 */
func (d *Dispatcher) dispatcher() {
  fmt.Println("开始执行 dispatcher -------------------")
  fmt.Println(len(Job_queue))
  i := 1
  for {
    select {
    // todo: 这里很重要， 这里是读取 Job_queue, 然后在request.go 那里的 run 方法是 把 新建的job写进 Job_queue, 这里才是真正把 读取到的job_queue 写进 worker 里面去
    case job := <-Job_queue:
      fmt.Println("开始分发 Job_queue 中的第 ", i, " 个job ------------")

      go func(job Job) {
        // 堵塞监听worker_pool的写入
        // 在 worker_pool 写入完毕之前， 这句逻辑代码一直堵塞
        // 尝试获取一个可用的 worker job channel
        // 阻塞直到有可用的 worker
        // 顺便定义了job_channel

        // todo: 因为根本就没有程序写入过 worker_pool, 所以这里会一直堵塞住  woerk_pool 这个channel

        // todo: 只要 worker Start() 还没处理完 worker_pool  的东西， 这里都会堵塞

        // todo: 堵塞读取 work_pool， 不断监听 worker_pool 的写入

        // todo: 关键理解点是， channel的特性， channel是 有东西在读的时候， 它不给其他东西写， 有东西在写的时候， 它不给其他东西读
        // todo: 这里这样写是为了原子性考虑， 有个东西正在读 work_pool channel, 那么就不给写入了， 这样就保证了数据的一致性
        job_channel := <-d.worker_pool

        // todo: 上面在堵塞channel， 不影响同线程的逻辑继续执行， 这里继续执行， 这里把 从 Job_queue 取得的job写进  job_channel, 写完为止
        // 分发任务到 worker job channel 中

        // todo: 把要处理的job 写进 job_channel, 写入的时候， 在 worker 的 Start 里面的 work_pool 就会读取 job_channel的数据， 因为 work_pool 的len 是3， 所以每次只能读取 3 个job， 每次只能处理3个job
        job_channel <- job

      }(job)      // 这个 job就是 case 上面获取的job

      i += 1
    }
  }
}








