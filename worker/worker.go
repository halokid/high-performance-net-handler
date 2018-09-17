package worker

import "fmt"


// fixme: 先定义一些处理的常量，这些常量优化时是要放在配置文件的
var (
  max_worker = 5
  max_queue = 2
)



// 首先应该定义主干结构体 worker, 这样有利于理清程序的思路， 其他的那些次要的结构体， job， payload等稍后再定义
// worker 储存要处理的任务
type Worker struct {
  work_pool       chan chan Job     // Job 缓冲池的缓冲池
  job_channel     chan Job          // fixme: 既然上面都定义了 work_pool, 这里不是多此一举吗？
  quit            chan bool
}


type PayLoad int

type Job struct {
  pay_load PayLoad
}


func (p PayLoad) Do() (err error) {
  fmt.Println("I am working Do", int(p))
  err = nil
  return err
}


func NewWorker (work_pool chan chan Job) Worker {
  return Worker{
    work_pool:      work_pool,
    job_channel:    make(chan Job),
    quit:           make(chan bool),
  }
}


// worker开始处理任务
func (w Worker) Start() {
  // 并行处理worker 的任务
  go func() {
    // 这种方式一般用来不断的监听channel的
    for {
      w.work_pool <- w.job_channel      // 是不是有点奇怪， 怎么把同一个结构体的变量赋予自己呢？

      select {
      // 在 job_channel 这个 channel 里面取出 job， job就是要处理的任务
      case job := <- w.job_channel:
        //  开始处理任务， 这里有两个方式来写这段代码
        //  1.  定义一个通用的处理函数， 来处理
        //  2.  定义一个 属于 job 结构体的函数来处理
        // 两者有什么区别呢？？ 一般都是用 2 来写， 为什么呢？  2 比 1 好吗？ 好在哪里
        // 这里先用 2 来写
        // fixme: 本身可以直接在 Job 的机构体来定义， 为什么还要在里面定义一个 pay_load 变量呢？
        // Do 函数为具体处理job的函数

        // 既然要搞成2 的方式， 那么 pay_load 就必须是一个结构体， 现在还不是， 所以要定义一个。。。 原来层层封装是这么来的， 靠
        if err := job.pay_load.Do(); err != nil {
          fmt.Printf("[ERROR] ---------@@@@@@@@@---------: %s", err.Error())
        }
      case <- w.quit:
        // 接收到要退出的信号， 则直接返回
        return
      }
    }
  }()
}


// 给 worker 传输一个停止的信号的函数
func (w Worker) Stop() {
  go func() {
    w.quit <- true
  }()
}



















