Phony := build

build:
	go build -o main main.go

run:
	cd cmd/default && go run main.go

validate-test:
	curl -X POST -H "Content-Type: application/json" -d '{"name":"김철수","age":25,"job":"designer"}' http://localhost:8080/health/validate

test:
	go test -v ./...