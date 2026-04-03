# ?? Multi-stage Build: Kucuk ve Guvenli Image
# AS builder asamasnda Go binary derlenir.
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/api

# AS final asamasnda sadece binary kopyalanr.
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
