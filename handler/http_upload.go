package handler

import (
  "net/http"
  "github.com/r00tjimmy/high-performance-net-handler/utils"
  "os"
  "io"
  "fmt"
  "strings"
  "time"
)

func HttpUploadHandle(w http.ResponseWriter, r *http.Request) {
  form_file, header, err := r.FormFile("uploadfile")
  utils.CheckErr(err)
  defer form_file.Close()

  //create save file
  dest_file, err := os.Create(utils.Http_upload_path + "/" + header.Filename)
  utils.CheckErr(err)
  defer dest_file.Close()

  // save file
  _, err = io.Copy(dest_file, form_file)
  utils.CheckErr(err)

  fmt.Fprintf(w, "upload success")
}



/**
upload file to data-time folder
 */
func HttpUploadDateTimeHandle(w http.ResponseWriter, r *http.Request) {
  // check path
  exists, folder_path := DateTimeFolderExists()
  if exists {
    // folder exists

  } else {
    // folder not exists, create
    err := os.Mkdir(folder_path, os.ModePerm)
    utils.CheckErr(err)
  }

}


/**
check date-time folder name
folder display like:
`-- 2018-09-11
    |-- 11-23
    |   |-- xxxxxxx.xml
    |   `-- yyyyyyyyyyyy.xml
    `-- 11-25
        |-- aaaaaaaaaaaa.xml
        `-- bbbbbbbbbbb.xml
3 directories, 8 files
 */
func DateTimeFolderExists() (bool, string) {
  date_time := time.Now().String()
  //fmt.Println(date_time)

  folder_sli := strings.Split(date_time, " ")
  //fmt.Println(folder_sli)

  date_folder := folder_sli[0]
  //fmt.Println(date_folder)

  time_folder := strings.Split(folder_sli[1], ":")[0] + "-" + strings.Split(folder_sli[1], ":")[1]
  //fmt.Println(time_folder)

  folder_path := date_folder + "/" + time_folder
  if utils.PathExists(date_folder + "/" + time_folder) {
    //fmt.Println("folder exists")
    return false, folder_path
  } else {
    //fmt.Println("folder not exists")
    return true, folder_path
  }
}










