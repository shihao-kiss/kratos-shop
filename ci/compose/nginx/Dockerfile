FROM registry.cn-hangzhou.aliyuncs.com/shihao-images/centos:7

MAINTAINER li_ning_@outlook.com

# 更换为阿里云镜像源
RUN cd /etc/yum.repos.d/ \
    && sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-* \
    && sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://mirrors.aliyun.com|g' /etc/yum.repos.d/CentOS-* \
    && yum clean all \
    && yum makecache

RUN yum install -y make zlib zlib-devel pcre pcre-devel gcc gcc-c++ openssl openssl-devel libevent libevent-devel perl unzip net-tools wget initscripts \
    && yum clean all && rm -rf /var/cache/yum/*

# 复制工具包
ADD ./compose/nginx/nginx-1.12.1.tar.gz /usr/local/src

# FastDFS 与 Nginx 集成
WORKDIR /usr/local/src/nginx-1.12.1
RUN ./configure --with-cc-opt=-Wno-error --with-http_ssl_module
RUN make && make install

# 配置 Nginx
ADD ./compose/nginx/nginx.conf /usr/local/nginx/conf

# 配置 Nginx 404
ADD ./compose/nginx/404.jpg /usr/local/nginx/html
ADD ./compose/nginx/404.html /usr/local/nginx/html

# SSL证书
ADD ./compose/nginx/https.key /usr/local/nginx
ADD ./compose/nginx/https.pem /usr/local/nginx

# 设置启动脚本
RUN mkdir -p /usr/local/bin
COPY ./compose/nginx/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod 755 /usr/local/bin/entrypoint.sh && \
    sed -i 's/\r$//' /usr/local/bin/entrypoint.sh

# 设置为上海时区
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/shanghai" >> /etc/timezone

ENTRYPOINT ["sh", "/usr/local/bin/entrypoint.sh"]

WORKDIR /
CMD ["/bin/bash"]