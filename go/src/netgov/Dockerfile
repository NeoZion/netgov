FROM golang:latest

RUN GOPROXY=https://goproxy.cn go get github.com/go-delve/delve/cmd/dlv
RUN GOPROXY=https://goproxy.cn go get github.com/gin-gonic/gin

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app

CMD ["dlv", "debug", "--headless", "--listen=:2345", "--log", "--api-version=2"]