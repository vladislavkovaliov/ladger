# ---------- build ----------
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN RUN swag init -g ./cmd/api/main.go && \
    sed -i 's/"host": "localhost:8080"/"host": "192.168.1.111:8080"/' docs/swagger.json


RUN swag init -g ./cmd/api/main.go 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/api

# ---------- runtime ----------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
