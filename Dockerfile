FROM golang:1.13.4-alpine as builder
WORKDIR /go/src/app
COPY ./src .
RUN go mod download && \
    go build -o main .

FROM alpine:3.7
RUN apk update && rm -rf /var/cache/apk/*
WORKDIR /root
COPY --from=builder /go/src/app/main .
CMD ["./main"]