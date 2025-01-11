# Detect the OS and set the executable suffix
ifeq ($(OS),Windows_NT)
EXE := .exe
else
EXE :=
endif

# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o main$(EXE) cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@docker compose up --build

# Shutdown DB container
docker-down:
	@docker compose down

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main$(EXE)

# Live Reload
watch:
	@if [ "$(OS)" = "Windows_NT" ]; then \
		powershell -ExecutionPolicy Bypass -Command "if (Get-Command air -ErrorAction SilentlyContinue) { \
			air; \
			Write-Output 'Watching...'; \
		} else { \
			Write-Output 'Installing air...'; \
			go install github.com/air-verse/air@latest; \
			air; \
			Write-Output 'Watching...'; \
		}"; \
	else \
		if command -v air >/dev/null 2>&1; then \
			air; \
		else \
			go install github.com/air-verse/air@latest; \
			air; \
		fi; \
	fi

.PHONY: all build run test clean watch docker-run docker-down itest