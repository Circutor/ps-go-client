
NAME=ps-go-client
GO=go


tidy:
	go mod tidy
	go mod vendor

# Linter.
lint:
	golangci-lint version
	golangci-lint linters
	golangci-lint run --fix

# Test
test:
	go test -coverprofile=profile.cov ./...
	go tool cover -func profile.cov
	go vet ./...
	gofmt -l .

build:
	go build -o $(NAME) ./cmd/main.go

