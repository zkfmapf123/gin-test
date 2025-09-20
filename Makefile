Phony := build

build:
	go build -o main main.go

run:
	cd cmd/default && go run main.go

test:
	go test -v ./...