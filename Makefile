default: all

all: compile

compile: deps test
	go build ./cmd/gerardd

test: deps
	go test ./gerard-core/tests/

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
