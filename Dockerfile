FROM golang:latest
LABEL authors="wang2"

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/kingsill/gin-example
COPY . $GOPATH/src/github.com/kingsill/gin-example
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin-example"]