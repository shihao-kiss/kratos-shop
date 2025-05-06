#!/bin/sh

# 检查 NATS 服务器的健康状态
# 使用 curl 访问 NATS 监控接口的健康检查端点
curl -f http://localhost:8222/healthz || exit 1 