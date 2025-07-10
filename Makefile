APP_NAME := bookstore-api
MAIN := ./cmd/server

run-dev: ## Run in development mode using go run
	go run $(MAIN)

run-release: ## Build and run compiled binary
	go build -o bin/$(APP_NAME) $(MAIN)
	./bin/$(APP_NAME)

test: ## Run all tests
	go test ./... -v

