FROM ubuntu:latest
FROM golang:1.26-alpine AS builder

# --- 1. build image => cook --- #
# SQLite CGO complier and GCC build (package)
RUN apk add --no-cache build-base

WORKDIR /workspace

# copy dependency docker server
# /app/go.mod , /app/go.sum 
COPY go.mod go.sum ./
RUN go mod download

# copy local file -> docker file 
COPY internal/handlers/ .
COPY app/ .

# CGO build
ENV CGO_ENABLED=0
RUN go build -o server .

# --- start image => eat --- #
FROM alpine:3.19
RUN apk add --no-cache sqlite-libs ca-certificates
# copy CGO build -> new dir /app
WORKDIR /app

# binary file build
COPY --from=builder /workspace/server ./server

RUN chmod +x server

CMD ["./server"]
