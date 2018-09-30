package handler

import (
  "net/http"
  "github.com/r00tjimmy/high-performance-net-handler/utils"
  "os"
  "io"
  "fmt"
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
  // read file
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
check date-time folder name
 */
func checkDateTimeFolder() {

}










