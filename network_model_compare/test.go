package main

func main() {
  //time.Sleep(5 * time.Second)
  //println("sleep end ....")

  xChan := make(chan int)
  xChan <- 1

  //yChan :=  make(chan  chan  int)

  y := <- xChan
  //y <- xChan
  println(y)
}
