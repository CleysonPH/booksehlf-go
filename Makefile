env=dev

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Done."

.PHONY: coverage
coverage:
	@echo "Running tests..."
	@go test -coverprofile=coverage/coverage.out ./...
	@echo "Done."
	@echo "Generating coverage report..."
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Done."

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf coverage
	@echo "Done."

.PHONY: run
run:
	@echo "Running..."
	@go run cmd/app/app.go --env=$(env)
	@echo "Done."
