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

test:
	golint -set_exit_status ./...
	go vet ./...
	errcheck ./...
	go test ./... -v -covermode=atomic

build: clean format test
	go build

.PHONY: clean format test build
