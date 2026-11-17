FROM golang:1.26-alpine AS builder
WORKDIR /build

RUN apk add --no-cache git

COPY go.mod go.sum ./
ENV GOPROXY=direct
RUN go mod download

COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/main ./app/src

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# copy binaru file 
COPY --from=builder /build/main .

EXPOSE 8080
CMD ["./main"]
