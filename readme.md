high performance net handler use golang
============================

this is a tool for handle net requests with high performance process

- support all protocol, like http, tcp, udp etc.
- Job\Worker network model
- use golang channel


### get
```
go get github.com/r00tjimmy/high-performance-net-handler

```

### sample HTTP handler
```go
package main

import (
  "github.com/r00tjimmy/high-performance-net-handler/worker"
)

var (
  max_worker = 3
  max_job = 10
  handle_type = "http"   // set network protocol type
)

func main() {
  // make the worker, listening work_pool channel
  dispatcher := worker.NewDispatcher(max_worker, handle_type)
  dispatcher.Run()

  // get requet
  request := worker.NewRequest(max_job, handle_type)
  request.Run()

}

```

### examples code build && test && run
```
cd examples

# just build
make build

after build, you can run with ./hpnh in current folder

# just test
make gotest

# auto build and run 
make all



```

