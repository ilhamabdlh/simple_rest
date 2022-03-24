build:
	go build -o cmd/service service.go

run: 
	go run ./cmd/service.go