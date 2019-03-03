package utils

import "fmt"

// some env setting
var (
  Version     =   "1.0"
  DebugFlag  =   true
  LogFlag    =   false
  LogFile    =   "../logs/hpnh.log"
)


// set handler setting
var (
  HttpUploadPath      =       "../uploads"
)

func TestUtils()  {
  fmt.Println("Test Utils")
}
