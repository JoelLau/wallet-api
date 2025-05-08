FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN  go mod download

COPY . ./
RUN  go build -o app ./cmd/app/main.go

FROM alpine:3.21 AS runner

WORKDIR /app
COPY --from=builder /app/app ./
EXPOSE 8080
ENTRYPOINT [ "./app" ]
