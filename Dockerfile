FROM golang:1.18-alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \

WORKDIR /tmp

# 将源码拷贝到容器中
COPY . .

# 将代码编译成二进制可执行文件app
RUN go build -o app .

WORKDIR /app
RUN cp /tmp/app .
EXPOSE 8080

ENTRYPOINT ["/app/app"]