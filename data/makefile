SERVICE_NAME = {{name}}
BINARY = $(SERVICE_NAME)
GOFILES := $(shell find . -type f -name '*.go')

# ===== BUILD =====

build:
	@echo ">> Building $(BINARY)"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/$(BINARY) .

run:
	@echo ">> Running $(BINARY)"
	go run .

dev:
	@echo ">> Starting in DEV mode (air hot-reload)"
	@if ! command -v air >/dev/null 2>&1; then \
		echo "air not found — installing..."; \
		go install github.com/cosmtrek/air@latest; \
	fi
	air

# ===== HOUSEKEEPING =====

tidy:
	@echo ">> go mod tidy"
	go mod tidy

lint:
	@echo ">> Running golangci-lint"
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "golangci-lint not found — installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run ./...

test:
	@echo ">> Running tests"
	go test ./... -v

# ===== DOCKER =====

docker:
	@echo ">> Building Docker image"
	docker build -t $(SERVICE_NAME):latest .

docker-run:
	@echo ">> Running Docker container"
	docker run --rm -p 5300:5300 $(SERVICE_NAME):latest

# ===== DOCKER COMPOSE =====

compose-up:
	@echo ">> Starting docker-compose"
	docker-compose up -d --build

compose-down:
	@echo ">> Stopping docker-compose"
	docker-compose down

# ===== CLEAN =====

clean:
	@echo ">> Cleaning build artifacts"
	rm -f ./bin/$(BINARY)

.PHONY: build run dev tidy lint test docker docker-run compose-up compose-down clean
