gotest:
	go test -coverprofile=coverage.out ./...

coverage:
	go tool cover -html="coverage.out"

dev:
	(cd ./cmd/server && go run main.go)
