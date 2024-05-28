.PHONY: build clean test run

build:
	@go build -o bin/e-commerce-golang-next cmd/main.go

test:
	@go test -v ./...

run:
	@./bin/e-commerce-golang-next

clean:
	rm -rf bin/
