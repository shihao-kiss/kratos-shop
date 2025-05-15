FROM registry.cn-hangzhou.aliyuncs.com/shihao-images/kafka:2.13-2.8.1

# 设置工作目录
WORKDIR /opt/kafka

# 设置环境变量
ENV KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"

# 暴露端口
EXPOSE 9092

# 启动命令
CMD ["start-kafka.sh"] 