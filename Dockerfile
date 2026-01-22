# ---------- build ----------
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go install github.com/swaggo/swag/cmd/swag@latest

# RUN swag init -g ./cmd/api/main.go 

# RUN if [ -f docs/swagger.json ]; then \
#       sed -i 's/"host": "localhost:8080"/"host": "192.168.1.111:8080"/' docs/swagger.json; \
#     fi

RUN cat docs/swagger.json 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/api

# ---------- runtime ----------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
