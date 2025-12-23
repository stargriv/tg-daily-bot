.PHONY: build run clean test install deploy help

# Binary name
BINARY=tg-daily-bot
# Build directory
BUILD_DIR=.

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-s -w"

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the bot binary
	@echo "Building $(BINARY)..."
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) -v

build-linux: ## Build for Linux x86_64 (64-bit - most common for VPS)
	@echo "Building $(BINARY) for Linux x86_64..."
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux -v

build-linux-32: ## Build for Linux x86 (32-bit)
	@echo "Building $(BINARY) for Linux x86 (32-bit)..."
	GOOS=linux GOARCH=386 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux-32 -v

build-arm: ## Build for ARM (Raspberry Pi, etc.)
	@echo "Building $(BINARY) for ARM..."
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-arm64 -v

run: build ## Build and run the bot
	@echo "Running $(BINARY)..."
	./$(BINARY)

clean: ## Remove build artifacts
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/$(BINARY)
	rm -f $(BUILD_DIR)/$(BINARY)-linux
	rm -f $(BUILD_DIR)/$(BINARY)-linux-32
	rm -f $(BUILD_DIR)/$(BINARY)-arm64

test: ## Run tests
	@echo "Running tests..."
	$(GOTEST) -v ./...

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

setup: ## Initial setup (copy .env.example to .env)
	@if [ ! -f .env ]; then \
		echo "Creating .env from .env.example..."; \
		cp .env.example .env; \
		echo "Please edit .env with your configuration"; \
	else \
		echo ".env already exists, skipping..."; \
	fi

install: build ## Install the bot to /usr/local/bin
	@echo "Installing $(BINARY) to /usr/local/bin..."
	sudo cp $(BUILD_DIR)/$(BINARY) /usr/local/bin/

uninstall: ## Uninstall the bot from /usr/local/bin
	@echo "Removing $(BINARY) from /usr/local/bin..."
	sudo rm -f /usr/local/bin/$(BINARY)

# VPS deployment helpers
deploy-systemd: ## Deploy as systemd service (requires sudo)
	@echo "Creating systemd service..."
	@echo "[Unit]" | sudo tee /etc/systemd/system/$(BINARY).service
	@echo "Description=Telegram Daily Bot" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "After=network.target" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "[Service]" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "Type=simple" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "User=$(USER)" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "WorkingDirectory=$(PWD)" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "ExecStart=$(PWD)/$(BINARY)" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "Restart=always" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "RestartSec=10" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "[Install]" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo "WantedBy=multi-user.target" | sudo tee -a /etc/systemd/system/$(BINARY).service
	@echo ""
	@echo "Reloading systemd..."
	sudo systemctl daemon-reload
	@echo ""
	@echo "Service created! Enable and start with:"
	@echo "  sudo systemctl enable $(BINARY)"
	@echo "  sudo systemctl start $(BINARY)"

start-service: ## Start the systemd service
	sudo systemctl start $(BINARY)

stop-service: ## Stop the systemd service
	sudo systemctl stop $(BINARY)

restart-service: ## Restart the systemd service
	sudo systemctl restart $(BINARY)

status-service: ## Check systemd service status
	sudo systemctl status $(BINARY)

logs: ## View systemd service logs
	sudo journalctl -u $(BINARY) -f

enable-service: ## Enable systemd service to start on boot
	sudo systemctl enable $(BINARY)

disable-service: ## Disable systemd service from starting on boot
	sudo systemctl disable $(BINARY)

# Development helpers
dev: ## Run in development mode (auto-restart on file changes requires entr)
	@command -v entr >/dev/null 2>&1 || { echo "entr is required for dev mode. Install with: brew install entr (macOS) or apt install entr (Linux)"; exit 1; }
	@echo "Running in development mode (auto-restart on changes)..."
	find . -name '*.go' | entr -r make run

fmt: ## Format Go code
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

lint: ## Run golangci-lint (requires golangci-lint installed)
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint is not installed. Install from https://golangci-lint.run/"; exit 1; }
	@echo "Running linter..."
	golangci-lint run

update-deps: ## Update all dependencies to latest versions
	@echo "Updating dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy

all: clean deps build test ## Clean, download deps, build and test
