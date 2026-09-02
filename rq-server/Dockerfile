# syntax=docker/dockerfile:1

FROM golang:1.26 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /rq-server

FROM golang:1.26 AS runtime

WORKDIR /app
COPY --from=builder /rq-server /app/rq-server
COPY config/config.non.yml /app/config/config.non.yml
COPY config/config.prod.yml /app/config/config.prod.yml
EXPOSE 8083
ENTRYPOINT ["/app/rq-server"]