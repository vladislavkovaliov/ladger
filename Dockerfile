# ---------- build ----------
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cat docs/swagger.json 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/api

CMD ["sh", "-c", "echo '--- ENV ---' && env && echo '--- APP ---' && ./app"]


# ---------- runtime ----------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["sh", "-c", "echo '--- ENV ---' && env && echo '--- APP ---' && ./app"]

CMD ["./app"]
