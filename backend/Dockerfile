FROM golang:1.14.2 AS builder
WORKDIR /go/src/shoppie
COPY . .
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.io,direct && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/shoppie/app .
CMD ["./app"]