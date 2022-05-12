## build: builds the command line tool to dist directory
build:
	@echo Building gogo...
	@go build -o ./dist/gogo ./cmd/...
	@echo Build Successfull.

## test: runs all tests
test:
	@cd ./pkg && go test -v

## coverage: displays test coverage
coverage:
	@go test -cover ./pkg

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./pkg && go tool cover -html=coverage.out

## install: installs executable as a global command to the machine
install: build
	@./scripts/install.sh