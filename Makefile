.PHONY: run
run:
	@go run cmd/services/main.go

.PHONY: migrate
migrate:
	@go run migrate/main.go

.PHONY: test
test:
	@go test ./... -v | grep -E 'FAIL|ok'

.PHONY: test-coverage
test-coverage:
	@go test ./... -cover -v | grep -vE 'no test files|\ttesting: warning: no tests to run'
