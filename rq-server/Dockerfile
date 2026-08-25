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
COPY config/config.prod.yaml /app/config/config.prod.yaml
EXPOSE 8083
ENTRYPOINT ["/app/rq-server"]