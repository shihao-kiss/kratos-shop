FROM registry.cn-hangzhou.aliyuncs.com/shihao-images/nats:2.10-alpine

# 安装 curl 用于健康检查
RUN apk add --no-cache curl

# 配置文件
COPY ./compose/nats/nats-server.conf /etc/nats/nats-server.conf
RUN chmod 644 /etc/nats/nats-server.conf

# 健康检查脚本
COPY ./compose/nats/healthcheck.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/healthcheck.sh

# 暴露端口
# 4222 是客户端连接端口
# 8222 是 HTTP 管理端口
# 6222 是集群通信端口
EXPOSE 4222 8222 6222

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD healthcheck.sh

# 启动命令
CMD ["nats-server", "--config", "/etc/nats/nats-server.conf"]

