FROM golang:1.26-alpine AS builder
WORKDIR /build

RUN apk add --no-cache git

COPY go.mod go.sum ./
ENV GOPROXY=direct

# main first module downalod 
RUN --mount=type=cache,target=/go/pkg/mod \
 go mod download 

COPY nginx/conf.d/*.conf /etc/nginx/conf.d/default.conf
COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/main ./app/src

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# copy binaru file 
COPY --from=builder /build/main .

EXPOSE 8080
CMD ["./main"]
