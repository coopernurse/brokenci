# Go Pretest Repository

This repository contains four branches with different CI issues for testing AI code review tools.

## Branches

1. **main** - All tests pass, code is properly formatted and linted
2. **test-failure** - Contains a failing unit test
3. **format-failure** - Code fails gofmt formatting check
4. **lint-failure** - Code fails golangci-lint checks

## Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/go-pretest-repo.git
cd go-pretest-repo

# Install development tools
make install-tools

# Run all checks
make quality
