FROM golang:1.17 AS builder

WORKDIR /go/src/
COPY . .
RUN GOOS=linux go build -o webapp

FROM debian:buster-slim

WORKDIR /app/
COPY --chown=0:0 --from=builder /go/src/webapp .

EXPOSE 3000
CMD ["./webapp"]