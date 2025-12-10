.PHONY: all test lint format check-fmt build clean verify-fix

# Default target runs all checks
all: verify-fix

# Run tests
test:
	@echo "Running tests..."
	go test ./... -v

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run ./...

# Check formatting
check-fmt:
	@echo "Checking formatting..."
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "Files need formatting:"; \
		gofmt -l .; \
		exit 1; \
	fi

# Format code
format:
	@echo "Formatting code..."
	gofmt -w .

# Build the project
build:
	@echo "Building..."
	go build -o target/pretest ./cmd/pretest

# Clean build artifacts
clean:
	@echo "Cleaning..."
	go clean
	rm -f pretest

quality: check-fmt lint test build
	@echo "✅ All checks passed!"
	@echo "✅ Ready to submit PR!"

# Install development tools
install-tools:
	@echo "Installing tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Show help
help:
	@echo "Available targets:"
	@echo "  make verify-fix   - Run all checks (for AI to verify fixes)"
	@echo "  make test         - Run tests"
	@echo "  make lint         - Run linter"
	@echo "  make check-fmt    - Check formatting"
	@echo "  make format       - Format code"
	@echo "  make build        - Build binary"
	@echo "  make install-tools - Install development tools"
	@echo "  make clean        - Clean build artifacts"
