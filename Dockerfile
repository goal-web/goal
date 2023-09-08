FROM golang:1.19 as builder
LABEL maintainer="qbhy <qbhy0715@qq.com>"

WORKDIR /app

COPY . /app
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPROXY=https://proxy.golang.com.cn,direct
RUN go build -ldflags="-s -w" -o app main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/app .

# run
ENTRYPOINT ["/app/app"]