default: all

all: compile

compile: deps test
	go build .

test: deps
	go test .

deps:
	git submodule update --init --recursive
