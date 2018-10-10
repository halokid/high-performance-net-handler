package test

import (
  "testing"
  "bytes"
  "mime/multipart"
  "github.com/r00tjimmy/high-performance-net-handler/utils"
  "github.com/r00tjimmy/high-performance-net-handler/handler"
  "os"
  "io"
  "net/http"
  "fmt"
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
  exists, folder_path := handler.DateTimeFolderExists()
  fmt.Println(exists, folder_path)

  err := handler.CreateDataFolder(folder_path)
  //err := handler.CreateDataFolder("aaaaaa")
  utils.CheckErr(err)

}












