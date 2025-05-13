FROM registry.cn-hangzhou.aliyuncs.com/shihao-images/consul:1.15

# 创建配置目录
RUN mkdir -p /consul/config

# 复制配置文件
COPY ./compose/consul/config.json /consul/config/config.json
# 修正权限
RUN chmod 755 /consul/config && chmod 644 /consul/config/config.json

# 健康检查脚本
COPY ./compose/consul/healthcheck.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/healthcheck.sh

# 暴露端口
# DNS - 8600
# HTTP - 8500
# HTTPS - disabled
# Server RPC - 8300
# LAN Serf - 8301
# WAN Serf - 8302
EXPOSE 8600 8500 8300 8301 8302

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD healthcheck.sh

# 启动命令
CMD ["consul", "agent", "-config-dir=/consul/config"] 