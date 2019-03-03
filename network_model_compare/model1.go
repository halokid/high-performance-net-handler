package main

import (
  "net/http"
  "time"
)



func main() {
  http.HandleFunc("/model1/", model1)
  err := http.ListenAndServe(":8089", nil)
  if err != nil {
    println("http serv ERROR")
  }
}

func model1(w http.ResponseWriter, r *http.Request) {
  println("model1 收到请求 ......")
  //time.Sleep(2 * time.Second)
  syncDo()
  w.Write([]byte("http请求等待同步逻辑处理完成后再返回"))
}

func syncDo() {
  println("model1 模拟同步处理请求， 处理逻辑花费的时间是2秒")
  time.Sleep(2 * time.Second)
}


