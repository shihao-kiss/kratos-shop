FROM registry.cn-hangzhou.aliyuncs.com/shihao-images/kafka:3.5
# 如有需要可COPY自定义脚本或配置

# 设置工作目录
WORKDIR /opt/kafka

# 复制自定义配置文件
COPY ./compose/kafka/server.properties /opt/kafka/config/server.properties

# 设置环境变量
ENV KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"

# 暴露端口
EXPOSE 9092

# 启动命令
CMD ["start-kafka.sh"] 