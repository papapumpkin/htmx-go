.PHONY: help build test clean run run-db generate-template shutdown

SHELL := /bin/bash

help:
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Compile the application
	@echo "Building the application..."

test: ## Run tests
	@echo "Running tests..."

clean: ## Remove temporary files and build artifacts
	@echo "Cleaning up..."

run: ## Start the application
	@echo "Starting the application..."
	@air -c .air.toml

run-db: ## Start the Postgres container with Podman
	@echo "Starting DB..."
	@podman compose up

generate-template: ## Generate templ code
	@echo "Generating templ code for frontend..."
	@templ generate

shutdown: ## Shutdown application
	@echo "Shutting down..."
	@podman compose down
