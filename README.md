# Wallet REST API

## Usage

1. generate files: `go generate ./...`
1. run: `go run cmd/app/main.go`

### Pre-requisites

1. [Go](https://go.dev/) >= 1.24 `brew install go@1.24`
2. [Goose](https://github.com/pressly/goose) `brew install goose`

### Commands

| Remarks           | Commands                     |
| ----------------- | ---------------------------- |
| Generate Files    | `go generate ./...`          |
| Run DB Migrations | `go run cmd/migrate/main.go` |
| Run HTTP Server   | `go run cmd/app-cli/main.go`     |
| Run Tests         | `go test -v -cover ./...`    |
| Run Linter        | `golangci-lint run`          |
| Fix Lint Issues   | `golangci-lint run --fix`    |
