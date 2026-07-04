BINARY_NAME := createpr
GO := go

.DEFAULT_GOAL := menu

# Colors
CYAN    := \033[36m
GREEN   := \033[32m
YELLOW  := \033[33m
DIM     := \033[2m
BOLD    := \033[1m
RESET   := \033[0m

menu:
	@printf "\n"
	@printf "$(BOLD)$(CYAN)╔══════════════════════════════════════════════════════════════╗$(RESET)\n"
	@printf "$(BOLD)$(CYAN)║                   createpr - Command Menu                    ║$(RESET)\n"
	@printf "$(BOLD)$(CYAN)╚══════════════════════════════════════════════════════════════╝$(RESET)\n"
	@printf "\n"
	@printf "  $(BOLD)$(GREEN)=== Development ===$(RESET)\n"
	@printf "   $(YELLOW)1)$(RESET)  make build            $(DIM)Build the binary$(RESET)\n"
	@printf "   $(YELLOW)2)$(RESET)  make install          $(DIM)Install globally via go install$(RESET)\n"
	@printf "   $(YELLOW)3)$(RESET)  make run              $(DIM)Run with go run$(RESET)\n"
	@printf "\n"
	@printf "  $(BOLD)$(GREEN)=== Testing ===$(RESET)\n"
	@printf "   $(YELLOW)4)$(RESET)  make test             $(DIM)Run all tests$(RESET)\n"
	@printf "   $(YELLOW)5)$(RESET)  make test-verbose     $(DIM)Run tests with verbose output$(RESET)\n"
	@printf "\n"
	@printf "  $(BOLD)$(GREEN)=== Code Quality ===$(RESET)\n"
	@printf "   $(YELLOW)6)$(RESET)  make fmt              $(DIM)Format code with gofmt$(RESET)\n"
	@printf "   $(YELLOW)7)$(RESET)  make lint             $(DIM)Check formatting + go vet$(RESET)\n"
	@printf "\n"
	@printf "  $(BOLD)$(GREEN)=== Maintenance ===$(RESET)\n"
	@printf "   $(YELLOW)8)$(RESET)  make clean            $(DIM)Remove build artifacts$(RESET)\n"
	@printf "   $(YELLOW)9)$(RESET)  make tidy             $(DIM)Tidy Go modules$(RESET)\n"
	@printf "\n"
	@read -p "  Enter choice: " choice; \
	case $$choice in \
		1) $(MAKE) build ;; \
		2) $(MAKE) install ;; \
		3) $(MAKE) run ;; \
		4) $(MAKE) test ;; \
		5) $(MAKE) test-verbose ;; \
		6) $(MAKE) fmt ;; \
		7) $(MAKE) lint ;; \
		8) $(MAKE) clean ;; \
		9) $(MAKE) tidy ;; \
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

fmt:
	gofmt -w .

fmt-check:
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "These files need formatting (run 'make fmt'):"; \
		gofmt -l .; \
		exit 1; \
	fi

lint: fmt-check
	$(GO) vet ./...

clean:
	rm -f $(BINARY_NAME)
	$(GO) clean

tidy:
	$(GO) mod tidy

help:
	@printf "\n"
	@printf "$(BOLD)Available commands:$(RESET)\n"
	@printf "\n"
	@printf "  $(CYAN)make build$(RESET)         Build the binary\n"
	@printf "  $(CYAN)make install$(RESET)       Install globally via go install\n"
	@printf "  $(CYAN)make run$(RESET)           Run with go run\n"
	@printf "  $(CYAN)make test$(RESET)          Run all tests\n"
	@printf "  $(CYAN)make test-verbose$(RESET)  Run tests with verbose output\n"
	@printf "  $(CYAN)make fmt$(RESET)           Format code with gofmt\n"
	@printf "  $(CYAN)make fmt-check$(RESET)     Check formatting (fails if unformatted)\n"
	@printf "  $(CYAN)make lint$(RESET)          Check formatting + go vet\n"
	@printf "  $(CYAN)make clean$(RESET)         Remove build artifacts\n"
	@printf "  $(CYAN)make tidy$(RESET)          Tidy Go modules\n"
	@printf "\n"

list: help

.PHONY: menu build install run test test-verbose fmt fmt-check lint clean tidy help list
