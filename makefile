
BINARY_NAME=Assistant
clean:
	@rm -rf bin
	@mkdir bin


build_linux: clean
	@echo "Building for Linux"
	@GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)_linux_amd64 cmd/Assistant/main.go

run: build_linux
	@echo "Running"
	@./bin/$(BINARY_NAME)_linux_amd64