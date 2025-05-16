FROM registry.cn-hangzhou.aliyuncs.com/shihao-images/zookeeper:3.8
# 如有需要可COPY自定义脚本或配置

# 设置工作目录
WORKDIR /opt/zookeeper

# 暴露端口
EXPOSE 2181

# 启动命令
CMD ["/opt/bitnami/zookeeper/bin/zkServer.sh", "start-foreground"] 