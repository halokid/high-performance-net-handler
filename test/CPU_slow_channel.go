package main

import(
  "fmt"
  "time"
  "runtime"
)

var quit chan int = make(chan int)

func loop(){
  for i:=0 ;i<10000; i++{
    fmt.Printf("%d\n", i)
  }

  quit <- 0
}

func main(){
  fmt.Println(runtime.NumCPU()) // 默认CPU核心数
  time.Sleep(2*time.Second)
  a := 100
  t0 := time.Now()      // 起点时间
  //runtime.GOMAXPROCS(1) // 设置执行使用的核心数

  for i:=1; i<= a; i++{
    go loop()
  }

  // fixme: 等待groutine完成的第一种方式：  如果没有消费掉 channel， 那么程序就不会等待
  for i:=0; i< a; i++{
    <- quit
  }

  endTime := time.Since(t0) // 耗时
  fmt.Println("运行时间", endTime)

}









