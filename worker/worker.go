package worker

import "fmt"

type Payload int

func (p Payload) Do() (err error)  {
  fmt.Println("payload Do() working at------", int(p), " job")
  err = nil
  return err
}

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

func NewWorker(work_pool chan chan Job) Worker  {
  return Worker{
    work_pool:      work_pool,
    job_channel:    make(chan Job),
    quit:           make(chan bool),
  }
}

func (w *Worker) Start() {
  go func() {
    for {
      // block listening w.job_channel
      // fixme: 在NewWorker() 这一步就已经初始化了 job_channel, 已经是有值的了，是一个内存位置， 默认是空值， 空值也是值， 只要有内存位置就可以是指针变量，就可以调用其方法

      // fixme： 依靠某个channel来做异步处理， 那么这个channel一定是要有读也有写的， 这个 job/work 模型的流程其实是这样的， 首先是 work.Start()方式, 把  job_channel 写进 work_pool， 然后 在 dispatcher方法里面读取 w.work_pool 那么就可以实现固定处理job的大小就是  work_pool 的 大小 3 个， 然后也可以实现异步响应了， 所谓的异步响应其实就是 一读一写， 一进一出

      // fixme: 首先是这里写入  w.work_pool
      w.work_pool <- w.job_channel

      select {
      case job := <- w.job_channel:
        if err := job.pay_load.Do(); err != nil {
          fmt.Println("[payload Do() ERROR] -------", err.Error())
        }

      case <- w.quit:
        return
      }
    }
  }()
}

func (w *Worker) Stop()  {
  go func() {
    w.quit <- true
  }()
}









