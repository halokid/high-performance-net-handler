package main

import (
  "./worker"
  "time"
  //"fmt"
)

// fixme: 先定义一些处理的常量，这些常量优化时是要放在配置文件的
// todo: 下面的意思是 ， 建 3 个worker， 去处理 20 个 job

/**
实际上 golang的 channel 跟我们普通的web程序的异步队列的原理是一样的， 比如我们普通web程序， 会先把请求
写进redis的队列里面去， 然后在后台有一个程序去处理 redis 队列里面的内容

但是在golang 来看， 我们完全可以把 接收请求---->  写进队列   -----> 后台处理做在同一个项目逻辑代码里面
因为这三个处理， 无非就是要异步处理才可以高效， 我们就用 channel来把 三步 逻辑分开就可以了

 */
var (
  max_worker = 3
  max_job = 20
)

func main() {
  // 总体逻辑就是， 收到请求 ----> 写入job 的pay_load -----> 通常 worker的start()来处理, 分到到Do()函数

  // 要先启动处理 request 的程序才可以， 而不是先接收request， 再启动处理， 这样比较合理
  // 分发请求, 处理payload， 也就是处理job
  dispatcher := worker.NewDispatcher(max_worker)
  dispatcher.Run()

  // fixme: 这个是测试数据， 生成job用的， 后面要写在test代码里, 实际上这里应该是一个循环监听的逻辑，不断接收请求， 然后生成job
  request := worker.NewRequest(max_job)
  request.Run()

  time.Sleep(time.Second * 10)
}



