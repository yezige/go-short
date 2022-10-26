FROM golang:1.18.3-alpine as build

ENV CGO_ENABLED=0 \
    GOPATH=/go \
    GOOS=linux \
    GOARCH=amd64

WORKDIR $GOPATH/short.liu.app

COPY . $GOPATH/short.liu.app

RUN go build -v -a -o short .

# 新增一个容器，用来运行应用
FROM alpine as run

RUN apk --no-cache add ca-certificates

WORKDIR /short.liu.app

RUN mkdir -p /var/log/short/

COPY ./*.yaml .

# 从 build 镜像中把/short 拷贝到当前目录
COPY --from=build /go/short.liu.app/short .

EXPOSE 8088

ENTRYPOINT ["./short"]

CMD ["-config", "./config_production.yaml"]
