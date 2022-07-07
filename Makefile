GENERATOR_BINARY_NAME="gosdk-codegen"
HELPER_SCRIPT_NAME="gencode.sh"
all: clean build generate format
generate:
	git submodule update --recursive --remote
	go run internal/main.go
	./$(HELPER_SCRIPT_NAME)
build:
	go build -o $(GENERATOR_BINARY_NAME) cmd/gosdk-codegen/main.go
format:
	gofmt -s -w  -l amzn/selling-partner-api-go-sdk/..
clean:
	go clean
	rm -f $(GENERATOR_BINARY_NAME)
	rm -f $(HELPER_SCRIPT_NAME)
	rm -rf internal/selling-partner-api-models-specs-v3/*.json
	rm -rf amzn/selling-partner-api-go-sdk/*