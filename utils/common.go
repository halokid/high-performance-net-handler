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
  if DebugFlag {
    log.Println(content)
  }

  if LogFlag {
    log_file_handle, err := os.OpenFile(LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
    CheckErr(err)
    log_file_handle.WriteString(content + "\n\n")
  }
}


/**
check os path exist or not
 */
func PathExists(path string) (bool) {
  _, err := os.Stat(path)
  if err == nil {
    return true
  }
  return false
}


/**
create folder
 */
func CreateFolder()  {

}






