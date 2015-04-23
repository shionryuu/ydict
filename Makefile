.PHONY: compile deps test

all: compile

compile: deps

deps:
	go get github.com/fatih/color
	go get github.com/ShionRyuu/ydict

test: compile
	go test -v ./...
