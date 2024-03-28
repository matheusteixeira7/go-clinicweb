test-unit:
	go test -v -failfast -coverprofile=coverage.out ./...

test-integration:
	go test -v -tags=integration -failfast -coverprofile=coverage.out ./...

coverage:
	go tool cover -html="coverage.out"

dev:
	(cd ./cmd/server && go run main.go)
