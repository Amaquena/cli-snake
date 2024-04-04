vendor:
	go mod tidy
	go mod vendor

run:
	go run ./...

build:
	go build -o clisnake
