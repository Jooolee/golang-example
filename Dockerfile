# 第一阶段：构建可执行文件
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件，下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制项目文件
COPY . .

# 构建可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 第二阶段：创建轻量级运行时镜像
FROM alpine:latest

# 安装必要的依赖
RUN apk --no-cache add ca-certificates
# 如果容器使用的是 Alpine Linux 镜像，需要使用 apk 包管理工具来安装 bash
RUN apk add --no-cache bash

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/main .

# 暴露服务监听的端口
EXPOSE 9000

# 运行应用程序
CMD ["./main"]

# docker build -t gin-service .
# docker run -p 9001:9000 --name gin-service gin-service 