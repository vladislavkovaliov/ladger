# ---------- build ----------
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cat docs/swagger.json 

ARG DATABASE_URL
RUN echo $DATABASE_URL

ENV DATABASE_URL=$DATABASE_URL
ENV PORT=$PORT
ENV JWT_SECRET=$JWT_SECRET
ENV JWT_EXPIRATION=$JWT_EXPIRATION

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/api

# ---------- runtime ----------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
