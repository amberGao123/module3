# buidler
FROM golang:1.18-alpine AS builder

# 配置编译环境
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 工作目录
WORKDIR /app
# 拷贝
COPY . /app/
# 编译二进制文件
RUN go build -v -o /bin/module3 ./main.go

# 基础镜像
FROM alpine
# 拷贝二进制文件
COPY --from=builder /bin/module3 /bin/module3
# 暴露的端口
EXPOSE 8080
# 指定容器启动程序
ENTRYPOINT [ "/bin/module3" ]
# 启动参数
CMD ["--addr=8080"]