FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN go build -o main .

FROM scratch AS final

WORKDIR /app
COPY --from=builder /app/main .

ENTRYPOINT ["/app/main"]
