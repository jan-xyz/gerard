default: all

all: compile

compile: deps test
	go build .

test: deps
	go test .

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
