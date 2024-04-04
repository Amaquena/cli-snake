vendor:
	go mod tidy
	go mod vendor

run:
	go run ./cmd/ ...

build:
	go build -o clisnake ./cmd
