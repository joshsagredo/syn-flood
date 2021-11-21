lint:
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

ineffassign:
	go get github.com/gordonklaus/ineffassign
	go mod vendor
	ineffassign ./...

test:
	sudo sysctl -w fs.file-max=10000000
	sudo go test ./... -v

build:
	go build -o bin/main cmd/syn-flood/main.go

run:
	sudo /usr/local/go/bin/go run cmd/syn-flood/main.go

cross-compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 cmd/syn-flood/main.go
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386 cmd/syn-flood/main.go
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 cmd/syn-flood/main.go
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 cmd/syn-flood/main.go
        # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/main-freebsd-amd64 cmd/syn-flood/main.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 cmd/syn-flood/main.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 cmd/syn-flood/main.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 cmd/syn-flood/main.go

upgrade-deps:
	go get -u ./...
	go mod tidy
	go mod vendor

all: test build run
