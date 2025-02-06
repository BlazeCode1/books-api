FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api .
FROM registry.trendyol.com/platform/base/image/appsec/chainguard/static/library:lib-20230201
COPY --from=builder /app/api /api
CMD ["./api"]