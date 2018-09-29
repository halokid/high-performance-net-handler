package utils

import (
  "log"
  "os"
)

func CheckErr(err error) {
  if err != nil {
    panic(err.Error())
  }
}

/**
output debug info & save log file
 */
func DebugLog(content string) {
  if debug_flag == true {
    log.Println(content)
  }

  if log_flag == true {
    log_file_handle, err := os.OpenFile()
  }
}
