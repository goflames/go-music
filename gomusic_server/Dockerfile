# 1. 使用官方的 golang 镜像作为基础镜像
FROM golang:1.21.3 AS build

# 2. 设置工作目录为 /app
WORKDIR /app

# 3. 将当前目录下的所有文件复制到容器的 /app 目录
COPY . .

# 下载依赖并编译应用程序
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# 使用一个更小的基础镜像来运行应用程序
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=build /app/main .

# 将 config/ 目录下的 config.yaml 文件复制到 /root/config 目录
COPY --from=build /app/config/config.yaml /root/config/


# 暴露服务的端口
EXPOSE 8888

# 运行可执行文件
CMD ["./main"]
