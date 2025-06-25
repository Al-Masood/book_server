#First Stage
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o book_server


#Second Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/book_server .

EXPOSE 3000

ENTRYPOINT ["./book_server"]
CMD ["serve", "-p=3000"]