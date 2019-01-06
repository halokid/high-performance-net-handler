package main

import (
  "fmt"
  "time"
)

/**
Job/Work 模式

注明的web服务器 nginx 用的也是这个网络模型
模型的逻辑原理：
1.  生成 N 个worker， 最多只能 N 个 worker 同时在处理job

2.  当有请求进来的时候， 生成 Job

3.  处理当有更多的job的时候， 暂时job是在等待 worker 处理的

这个模型的优点就是  worker 不会因为 job 太多而 因为阻塞处理 job 耗光了计算机资源， 而影响自身
应该处理的 worker 逻辑
 */



// worker本身是应该有动作， 比如run等，所以worker应该定义为一个结构体，而不是像下面定义为一个变量
/**
var (
  //最大的worker数量为3
  workerPool = make(chan [3]int)
)
**/

// 抽象化， job 也应该为一个结构体
type Job struct {
  payLoad     int
}

var JobQueue chan Job
func init() {
  JobQueue = make(chan Job, 1)
}

func (job *Job) doAction() {
  println("Job do action")
}

// worker 是结构体， 生成应该放进内存里面
type Wokrer struct {
  // 代表正在处理的job， 当job在处理的时候， 就放进这个channel
  workPool      chan chan Job      // 一个 channel类型的变量， channel里面放的是channel类型的变量
  jobChannel    chan Job
  quit          chan bool
}

// 有结构体， 通常都需要一个New方法
func newWorker(workPool chan chan Job) Wokrer {
  return Wokrer{
    workPool:       workPool,
    jobChannel:     make(chan Job),
    quit:           make(chan bool),
  }
}

func (w *Wokrer) start() {
  println("worker start ----------------\n\n")
  go func() {
    println("worker start go func  ----------------\n\n")
    for {
      w.workPool <- w.jobChannel   // 首先 wokerPool 读取， 获取要处理的job, 这个时候 workPool 的数量加一， 一直加到 3 为止， 就停止读取， 那么就没有下面的 select 行为
      select {
      case job := <- w.jobChannel:   // 读取 job, 和上面的 w.workdPool 读取到的是同样的值来的
        job.doAction()               // 具体 job 要做的事
      }
    }
  }()
}


func main() {
  // --------------  生成Worker START -----------------------------

  println("worker逻辑开始 ............\n\n")
  // 生成 3 个 worker 并 启动worker
  workPool := make(chan chan Job, 3)
  // 同一个 workPool 做处理控制， 然后靠 jobChannel 不断的写进 workPool
  for i := 0; i < 3; i++ {
    fmt.Printf("开始循环生成worker--------%d\n", i)
    worker := newWorker(workPool)
    // 开始 worker
    worker.start()
  }

  // 如果 没有init 函数先生成一个  job, 那么跑到这个函数这里就会阻塞， 当然可以把 生成 job 的逻辑提前到这个函数之前
  // 但是真实场景的 job生成的逻辑肯定在这个之前， 所以先用 init 函数， 模拟盛恒一个 job
  go dispatcher(workPool)
  // 同一个线程， 还需要有获得job的逻辑,  获得job 要靠下面的 job 生成， 然后写进 jobQueue 队列里面

  // --------------  生成Worker END -----------------------------



  // --------------  生成Job START -----------------------------
  // 生成 10 个 job， 先不用HTTP请求来生成
  for i := 0; i < 10; i++ {
    jobIndex := 1
    job := Job{payLoad:     jobIndex}
    JobQueue <- job
  }

  time.Sleep(time.Second * 10)
  // --------------  生成Job END -----------------------------
}



func dispatcher(workerPool chan chan Job) {
  for {
    select {
    case job := <- JobQueue:
      go func(job Job) {
        //job_channel := <- d.work_pool
        // 如果 workerPool 没有数据， 这里会阻塞住， 一直等待
        //fixme: 当处理请求多的时候， 其实就是会阻碍在这里的， 但是这种模型也正是因为阻塞在这里， 而不会影响worker的处理逻辑
        job_channel := <- workerPool
        //_ := <- d.work_pool

        // 把 job channel 赋值给 job_channel
        job_channel <- job
      }(job)
    }
  }
}





