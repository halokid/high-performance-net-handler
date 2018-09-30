package utils

import "fmt"

// some env setting
var (
  version     =   "1.0"
  debug_flag  =   true
  log_flag    =   false
  log_file    =   "../logs/hpnh.log"
)


// set handler setting
var (
  Http_upload_path      =       "../uploads"
)

func TestUtils()  {
  fmt.Println("Test Utils")
}
