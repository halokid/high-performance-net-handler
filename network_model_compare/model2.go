package main

/**
go channel 的异步模式
 */
import (
  "net/http"
  "time"
)

func model2(w http.ResponseWriter, r *http.Request) {
  println("model2 收到请求 ......")
  go func() {
    asyncModel2Do()
  }()

  w.Write([]byte("http请求不用实时等待逻辑处理完成， 分发给了异步函数去处理，可以马上返回"))
}

func asyncModel2Do()  {
  println("model2 模拟异步处理请求， 处理逻辑花费的时间是2秒")
  time.Sleep(2 * time.Second)
}


func main() {
  http.HandleFunc("/model2/", model2)
  err := http.ListenAndServe(":8089", nil)
  if err != nil {
    println("http serv ERROR")
  }
}
