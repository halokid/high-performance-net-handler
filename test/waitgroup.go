package main

import (
  "sync"
  "time"
  "fmt"
)

func main() {
  var wg sync.WaitGroup

  for i := 0; i < 5; i++ {
    wg.Add(1)

    go func(n int) {
      // 下面两句效果是一样的
      //defer wg.Add(-1)
      defer wg.Done()
      EchoNumber(n)
    }(i)
  }

  // fixme: 等待所有 groutine完成的第二种方式
  wg.Wait()
}

func EchoNumber(i int)  {
  time.Sleep(3e9)
  fmt.Println(i)
}









