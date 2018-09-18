all:    build
  
clean:
	rm -f coverage.txt
	
format:
	gofmt -w .
	gofmt -w gzip/
	gofmt -w snappy/
	gofmt -w xz/

test:
	golint -set_exit_status .
	golint -set_exit_status gzip/
	golint -set_exit_status snappy/
	golint -set_exit_status xz/
	go vet ./...
	errcheck ./...
	go test ./... -v -covermode=atomic

build: clean format test
	go build

.PHONY: clean format test build
