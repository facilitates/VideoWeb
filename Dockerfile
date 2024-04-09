# 使用官方Go镜像作为构建环境
FROM golang:alpine AS builder

# 设置工作目录
WORKDIR /build

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o videoweb .

# 使用alpine作为基础镜像
FROM alpine:latest

# 添加时区支持
RUN apk --no-cache add ca-certificates tzdata
ENV TZ Asia/Shanghai

# 设置工作目录
WORKDIR /root/

# 从构建器镜像中复制二进制文件
COPY --from=builder /build/videoweb .

# 安装MySQL
RUN apk add --no-cache mysql mysql-client
# 初始化MySQL数据库
RUN mysql_install_db --user=mysql --ldata=/var/lib/mysql

# 安装Redis
RUN apk add --no-cache redis

# 暴露端口
EXPOSE 3000 3306 6379

# 运行应用程序
CMD ["./videoweb"]
