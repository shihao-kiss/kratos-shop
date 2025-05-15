FROM registry.cn-hangzhou.aliyuncs.com/shihao-images/kafka:2.13-2.8.1

# 设置工作目录
WORKDIR /opt/kafka

# 复制配置文件
COPY compose/kafka/server.properties /opt/kafka/config/server.properties

# 设置环境变量
ENV KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"

# 暴露端口
EXPOSE 9092

# 启动命令
CMD ["start-kafka.sh"] 