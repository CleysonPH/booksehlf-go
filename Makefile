env=dev
dsn="mysql://root:root@tcp(localhost:3306)/bookshelf"

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Done."

.PHONY: coverage
coverage:
	@echo "Running tests..."
	@mkdir -p coverage
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

.PHONY: makemigration
makemigration:
	@echo "Creating migration..."
	@migrate create -ext sql -dir migrations -seq $(name)
	@echo "Done."

.PHONY: migrateup
migrateup:
	@echo "Migrating up..."
	@migrate -path migrations -database $(dsn) -verbose up
	@echo "Done."

.PHONY: migratedown
migratedown:
	@echo "Migrating down..."
	@migrate -path migrations -database $(dsn) -verbose down
	@echo "Done."