default: all

all: test build lint

build:
	go build -v .

test:
	go test -v .

lint:
	go get -u golang.org/x/lint/golint
	golint .
