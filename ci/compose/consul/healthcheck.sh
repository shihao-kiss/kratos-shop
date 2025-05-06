#!/bin/sh

# 检查 Consul 服务器的健康状态
# 使用 consul 内置的健康检查命令
consul members || exit 1 