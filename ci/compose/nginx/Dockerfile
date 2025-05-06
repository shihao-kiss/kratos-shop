FROM centos:7

MAINTAINER li_ning_@outlook.com

RUN yum install -y make zlib zlib-devel pcre pcre-devel gcc gcc-c++ openssl openssl-devel libevent libevent-devel perl unzip net-tools wget initscripts \
    && yum clean  all && rm -rf /var/cache/yum/*

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

COPY ./compose/nginx/entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

# 设置为上海时区
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/shanghai" >> /etc/timezone

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

WORKDIR /
CMD ["/bin/bash"]