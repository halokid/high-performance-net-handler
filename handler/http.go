package handler

import (
  "fmt"
  "net/http"
)


// process frontend sample
func HttpFrontSample(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "http frontend sample")
}


// process backend sample
func HttpBackendSample() {
  fmt.Println("http handle process in backend\n\n")
}

func HttpProcess() {
  // TODO: do your backend work here
  HttpBackendSample()
}



