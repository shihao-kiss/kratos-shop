FROM registry.cn-hangzhou.aliyuncs.com/shihao-images/mysql:5.7

RUN rm -f /etc/mysql/my.cnf
COPY ./compose/mysql/mysql.cnf /etc/mysql/my.cnf

# 拷贝数据库初始化sql脚本
COPY ./compose/mysql/init.sql /docker-entrypoint-initdb.d/init.sql

# 设置为上海时区
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/shanghai" >> /etc/timezone
