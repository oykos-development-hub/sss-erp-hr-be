BINARY_NAME=hrApp

build:
	@echo "Building HR app..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "HR app built!"

run: build
	@echo "Starting HR app..."
	@./tmp/${BINARY_NAME}
	@echo "HR app started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping HR app..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped HR app!"

restart: stop start