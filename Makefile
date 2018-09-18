all:    build
  
clean:
	rm -f coverage.txt
	
format:
	gofmt -w .
	gofmt -w gzip/
	gofmt -w lzw/
	gofmt -w snappy/
	gofmt -w xz/
	gofmt -w zlib/

deps:
	go get -u golang.org/x/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -t -v ./...

test:
	go vet ./...
	golint -set_exit_status ./...
	errcheck ./...
	go test ./... -v -covermode=atomic

build: clean format test
	go build

.PHONY: clean format test build
