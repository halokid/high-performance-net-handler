package utils

import "fmt"

// some env setting
var (
  Version     =   "1.0"
  Debug_flag  =   true
  Log_flag    =   false
  Log_file    =   "../logs/hpnh.log"
)


// set handler setting
var (
  Http_upload_path      =       "../uploads"
)

func TestUtils()  {
  fmt.Println("Test Utils")
}
