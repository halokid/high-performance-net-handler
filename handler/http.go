package handler

import (
  "fmt"
  "net/http"
)

func HttpHandle(w http.ResponseWriter, r *http.Request) {
  fmt.Println("HTTP handle start -------- ")
  fmt.Fprintf(w,"handle http request")
}
