
BINARY_NAME=assistant
clean:
	@rm -rf bin
	@mkdir bin

build:
	@echo "Building for local OS"
	@make build_linux

build_linux: clean
	@echo "Building for Linux"
	@GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)_linux_amd64 cmd/Assistant/main.go

run: build_linux
	@echo "Running"
	@./bin/$(BINARY_NAME)_linux_amd64