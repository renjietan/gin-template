SHELL=/usr/bin/env bash
NAME := gin-template

all: linux-amd64 linux-arm64 windows-amd64 windows-arm64

linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/$(NAME)-linux-amd64 main.go
.PHONY: linux-amd64

linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=7 go build -ldflags "-s -w" -o bin/$(NAME)-linux-arm64 main.go
.PHONY: linux-arm64

windows-amd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/$(NAME)-windows-amd64.exe main.go
.PHONY: windows-amd64

windows-arm64:
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -o bin/$(NAME)-windows-arm64.exe main.go
.PHONY: windows-arm64

clean:
	rm -f bin/$(NAME)-*
.PHONY: clean