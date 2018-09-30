package test

import (
  "testing"
  "bytes"
  "mime/multipart"
  "github.com/r00tjimmy/high-performance-net-handler/utils"
  "os"
  "io"
  "net/http"
  "time"
  "fmt"
  "strings"
)

func TestHttpUpload(t *testing.T) {
  // make form
  buf := new(bytes.Buffer)
  writer := multipart.NewWriter(buf)
  form_file, err := writer.CreateFormFile("uploadfile", "upload_test.txt")
  utils.CheckErr(err)

  // get data from file, write to form
  src_file, err := os.Open("upload_test.txt")
  utils.CheckErr(err)
  defer src_file.Close()
  _, err = io.Copy(form_file, src_file)

  // send form
  content_type := writer.FormDataContentType()
  writer.Close()
  _, err = http.Post("http://127.0.0.1:8089/hpnh_upload", content_type, buf)
  utils.CheckErr(err)
}


func TestGetDateTimeFolder(t *testing.T) {
  date_time := time.Now().String()
  fmt.Println(date_time)

  folder_sli := strings.Split(date_time, " ")
  fmt.Println(folder_sli)

  date_folder := folder_sli[0]
  fmt.Println(date_folder)

  time_folder := strings.Split(folder_sli[1], ":")[0] + "-" + strings.Split(folder_sli[1], ":")[1]
  fmt.Println(time_folder)

  if utils.PathExists(date_folder + "\\" + time_folder) {
    fmt.Println("folder exists")
  } else {
    fmt.Println("folder not exists")
  }
}












