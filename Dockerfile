# Building stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app/

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o confluence-payment


# Running stage
FROM alpine:latest

WORKDIR /root/

# Copying the binary
COPY --from=builder /app/confluence-payment . 
# Copying the .env file
# COPY --from=builder /app/.env .

EXPOSE 8001

CMD ["./confluence-payment"]