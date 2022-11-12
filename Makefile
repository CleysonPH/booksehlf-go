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