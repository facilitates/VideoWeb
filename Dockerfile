# 使用golang:alpine作为编译环境
FROM golang:alpine AS builder

# 设置工作目录
WORKDIR /build

# 复制go.mod和go.sum到工作目录
COPY go.mod go.sum ./

# 下载依赖包
RUN go mod download

# 复制当前目录到工作目录
COPY . .

# 编译Go应用程序，禁用CGO，指定目标操作系统为linux，输出可执行文件名为videoweb
RUN CGO_ENABLED=0 GOOS=linux go build -o videoweb .

# 最终使用的镜像
FROM alpine:latest

# 设置时区为上海
ENV TZ=Asia/Shanghai

# 设置工作目录为/root
WORKDIR /root

# 从builder阶段复制编译好的应用程序到当前镜像
COPY --from=builder /build/. .

EXPOSE 3000

CMD ["./videoweb"]
