build:
	go build && mv high-performance-net-handler	 hpnh

gotest:
	cd ./test &&  go test && cd ../ 

run:
	./hpnh	

all: build run


