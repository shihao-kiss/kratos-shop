FROM registry.cn-hangzhou.aliyuncs.com/shihao-images/zookeeper:latest

# 设置工作目录
WORKDIR /opt/zookeeper

# 复制配置文件
COPY compose/kafka/zookeeper/zoo.cfg /conf/zoo.cfg

# 暴露端口
EXPOSE 2181

# 启动命令
CMD ["/opt/zookeeper-3.4.13/bin/zkServer.sh", "start-foreground"] 