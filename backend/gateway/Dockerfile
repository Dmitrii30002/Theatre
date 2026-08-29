FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /gateway ./cmd

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /gateway /app/gateway
COPY config /app/config
EXPOSE 8080
ENTRYPOINT ["/app/gateway"]
