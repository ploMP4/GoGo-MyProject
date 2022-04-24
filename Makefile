## build: builds the command line tool to dist directory
build:
	@echo Building gogo...
	@go build -o ./dist/gogo .
	@echo Build Successfull.

## test: runs all tests
test:
	@cd ./cmd && go test -v

## coverage: displays test coverage
coverage:
	@go test -cover ./cmd

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./cmd && go tool cover -html=coverage.out