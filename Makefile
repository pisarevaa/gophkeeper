.PHONY: help unit_test integration_test e2e_test test lint coverage_report

help:
	cat Makefile

test:
	go clean -testcache && go test -v ./internal/...

mock_generate:
	mockgen -source=internal/server/storage/db/types.go -destination=internal/server/mocks/storage.go -package=mock AuthStorage,KeeperStorage
	mockgen -source=internal/server/storage/minio/types.go -destination=internal/server/mocks/minio.go -package=mock MinioStorage
	mockgen -source=internal/server/service/auth/types.go -destination=internal/server/mocks/auth.go -package=mock AuthServicer
	mockgen -source=internal/server/service/keeper/types.go -destination=internal/server/mocks/keeper.go -package=mock KeeperServicer

swag_generate:
	swag init --dir cmd/server,internal

swag_format:
	swag fmt

lint:
	go fmt ./...
	find . -name '*.go' -exec goimports -w {} +
	find . -name '*.go' -exec golines -w {} -m 120 \;
	golangci-lint run ./...

run_server:
	go run ./cmd/server

run_agent:
	go run ./cmd/agent

up_database_and_minio:
	docker compose -f docker-compose.yml up -d

generate_keys:
	go run ./cmd/generate_keys

coverage_report:
	go test -coverpkg=./... -count=1 -coverprofile=.coverage.out ./...
	go tool cover -html .coverage.out -o .coverage.html