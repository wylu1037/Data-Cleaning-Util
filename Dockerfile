FROM golang:1.18-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

# 将源码拷贝到容器中
COPY . .

# 将代码编译成二进制可执行文件app
RUN go build -o clear

FROM scratch

WORKDIR /app

COPY --from=builder /build/clear .
COPY --from=builder /build/config/application.ini ./config/

ENTRYPOINT ["/app/clear"]