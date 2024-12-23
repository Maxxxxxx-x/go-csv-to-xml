#!make

.DEFAULT_GOAL := all

tidy:
	go mod tidy
	go fmt ./...

all:
	go build -o=./converter ./main.go
