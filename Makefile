help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## build: builds the command line tool to dist directory for amd64-linux
build:
	@echo [Linux] Building gogo...
	@GOOS=linux GOARCH=amd64 go build -o ./dist/gogo-linux-amd64 ./cmd/...
	@echo Build Successfull.

## build: builds the command line tool to dist directory for amd64-windows
build-windows:
	@echo [Windows] Building gogo...
	@GOOS=windows GOARCH=amd64 go build -o ./dist/gogo-windows-amd64 ./cmd/...
	@echo Build Successfull.

## build: builds the command line tool to dist directory for amd64-macos
build-mac:
	@echo [macOS] Building gogo...
	@GOOS=darwin GOARCH=amd64 go build -o ./dist/gogo-macos-amd64 ./cmd/...
	@echo Build Successfull.

## build: builds the command line tool to dist directory for all amd64-(macos, windows, linux)
build-all: build build-windows build-mac

## test: runs all tests
test:
	@cd ./internal && go test -v

## coverage: displays test coverage
coverage:
	@go test -cover ./internal

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./internal && go tool cover -html=coverage.out

## install: installs executable as a global command to the machine
install: 
	@./scripts/install.sh

## doc: start documentation dev environment
doc:
	@cd docs && npm start
