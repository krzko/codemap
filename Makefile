.SILENT:

# Colours for logging
BLUE=\033[0;34m
GREEN=\033[0;32m
NC=\033[0m

# Platform emojis
LINUX_EMOJI=🐧
DARWIN_EMOJI=🍎
WINDOWS_EMOJI=🪟
AMD64_EMOJI=💻
ARM64_EMOJI=📱

# Logging functions
define log
	@echo "${BLUE}[$(shell date '+%Y-%m-%d %H:%M:%S')] ${GREEN}INFO${NC}  $(1)"
endef

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOLINT=golangci-lint

# Build settings
BINARY_NAME=codemap
MAIN_PACKAGE=./cmd/codemap
BUILD_DIR=./build
PLATFORMS=linux darwin windows
ARCHITECTURES=amd64 arm64

# Version information
VERSION=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

.PHONY: all build clean test coverage lint deps tidy run help

all: build

build: ## Build for all platforms
	$(call log,"🚀 Starting build process")
	@mkdir -p $(BUILD_DIR)
	@for os in $(PLATFORMS); do \
		for arch in $(ARCHITECTURES); do \
			platform_emoji=""; \
			arch_emoji=""; \
			case $$os in \
				linux) platform_emoji="$(LINUX_EMOJI)";; \
				darwin) platform_emoji="$(DARWIN_EMOJI)";; \
				windows) platform_emoji="$(WINDOWS_EMOJI)";; \
			esac; \
			case $$arch in \
				amd64) arch_emoji="$(AMD64_EMOJI)";; \
				arm64) arch_emoji="$(ARM64_EMOJI)";; \
			esac; \
			echo "${BLUE}[$(shell date '+%Y-%m-%d %H:%M:%S')] ${GREEN}INFO${NC}  Building $$platform_emoji $$os/$$arch_emoji $$arch"; \
			output=$(BUILD_DIR)/$(BINARY_NAME)-$$os-$$arch; \
			if [ $$os = "windows" ]; then output=$$output.exe; fi; \
			GOOS=$$os GOARCH=$$arch $(GOBUILD) $(LDFLAGS) -o $$output $(MAIN_PACKAGE); \
			echo "${BLUE}[$(shell date '+%Y-%m-%d %H:%M:%S')] ${GREEN}INFO${NC}  Completed $$platform_emoji $$os/$$arch_emoji $$arch"; \
		done; \
	done
	$(call log,"✨ Build process complete")

clean: ## Remove build artifacts
	$(call log,"🧹 Cleaning artifacts")
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	$(call log,"✨ Clean complete")

test: ## Run tests
	$(call log,"🧪 Starting tests")
	@$(GOTEST) -v ./...
	$(call log,"✅ Tests complete")

coverage: ## Run tests with coverage
	$(call log,"📊 Starting coverage tests")
	@$(GOTEST) -v -coverprofile=coverage.out ./...
	@$(GOCMD) tool cover -html=coverage.out
	$(call log,"📈 Coverage complete")

lint: ## Run linter
	$(call log,"🔍 Starting linter")
	@$(GOLINT) run
	$(call log,"✨ Linting complete")

deps: ## Download dependencies
	$(call log,"📦 Downloading dependencies")
	@$(GOGET) -v -t -d ./...
	$(call log,"✨ Dependencies complete")

tidy: ## Tidy and verify dependencies
	$(call log,"🧹 Tidying dependencies")
	@$(GOMOD) tidy
	@$(GOMOD) verify
	$(call log,"✨ Tidy complete")

run: build ## Run the application
	$(call log,"🚀 Starting application")
	@$(BUILD_DIR)/$(BINARY_NAME)-$(shell go env GOOS)-$(shell go env GOARCH)$(if $(findstring windows,$(shell go env GOOS)),.exe,)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help