build:
	@go build -o bin/go-minibank
run:build
	@./bin/go-minibank
test:
	@go test -v ./...