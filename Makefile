BINARY_NAME := createpr
GO := go

.DEFAULT_GOAL := menu

menu:
	@echo "╔═══════════════════════════════════════════════════╗"
	@echo "║          createpr — Command Menu                  ║"
	@echo "╚═══════════════════════════════════════════════════╝"
	@echo ""
	@echo "  === Development ==="
	@echo "  1) Build binary"
	@echo "  2) Install globally"
	@echo "  3) Run (go run)"
	@echo ""
	@echo "  === Testing ==="
	@echo "  4) Run all tests"
	@echo "  5) Run tests (verbose)"
	@echo "  6) Run linter"
	@echo ""
	@echo "  === Maintenance ==="
	@echo "  7) Clean build artifacts"
	@echo "  8) Tidy modules"
	@echo ""
	@read -p "Enter choice: " choice; \
	case $$choice in \
		1) $(MAKE) build ;; \
		2) $(MAKE) install ;; \
		3) $(MAKE) run ;; \
		4) $(MAKE) test ;; \
		5) $(MAKE) test-verbose ;; \
		6) $(MAKE) lint ;; \
		7) $(MAKE) clean ;; \
		8) $(MAKE) tidy ;; \
		*) echo "Invalid choice" ;; \
	esac

build:
	$(GO) build -o $(BINARY_NAME) .

install:
	$(GO) install .

run:
	$(GO) run .

test:
	$(GO) test ./...

test-verbose:
	$(GO) test -v ./...

lint:
	$(GO) vet ./...

clean:
	rm -f $(BINARY_NAME)
	$(GO) clean

tidy:
	$(GO) mod tidy

help:
	@echo "Available commands:"
	@echo "  make build         - Build the binary"
	@echo "  make install       - Install globally via go install"
	@echo "  make run           - Run with go run"
	@echo "  make test          - Run all tests"
	@echo "  make test-verbose  - Run tests with verbose output"
	@echo "  make lint          - Run go vet"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make tidy          - Tidy Go modules"

list: help

.PHONY: menu build install run test test-verbose lint clean tidy help list
