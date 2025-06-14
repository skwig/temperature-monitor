FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=1

# gcc etc
RUN apk add build-base

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:3.22 AS final

WORKDIR /app
COPY --from=builder /app/main .

ENTRYPOINT ["/app/main"]
