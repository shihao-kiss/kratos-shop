#!/bin/bash

# 脚本用法说明
# sh update.sh                                                # 更新所有容器
# sh update.sh servername                                     # 更新指定容器

# 环境变量设置
RUN_ENV_TYPE="product"                                        # 设置为生产环境

# 输出当前服务信息
echo -e "\033[32m Cur Service: $1 ... \033[0m"               # 绿色文字显示当前服务名

# sh update.sh 默认 升级 yml 配置中所有容器
# sh update.sh servername 只升级 servername

if [[ -n "$1" ]]; then                                       # 判断是否指定了服务名
    echo -e "\033[32m ${RUN_ENV_TYPE} $1 ... \033[0m"       # 显示环境类型和服务名

    # 停止指定容器
    echo -e "\033[32m Stopping container $1 ... \033[0m"     # 提示停止容器
    docker-compose -f ${RUN_ENV_TYPE}_compose.yml stop $1    # 停止指定的容器
    sleep 3s                                                 # 等待3秒确保停止完成

    # 删除指定容器
    echo -e "\033[32m Remove container $1 ... \033[0m"       # 提示删除容器
    docker-compose -f ${RUN_ENV_TYPE}_compose.yml rm -f $1   # 强制删除指定容器
    sleep 3s                                                 # 等待3秒确保删除完成

    # 更新代码
    echo -e "\033[32m Pulling the latest code from the remote... \033[0m"  # 提示拉取代码
    git pull origin main                                    # 从main分支拉取最新代码
    sleep 3s                                                 # 等待3秒确保拉取完成

    # 设置超时时间
    echo -e "\033[32m Set environment variables(120s)... \033[0m"  # 提示设置超时
    export DOCKER_CLIENT_TIMEOUT=120                         # Docker客户端超时设置
    export COMPOSE_HTTP_TIMEOUT=120                          # Compose HTTP超时设置

    # 构建镜像
    echo -e "\033[32m $1 Images is being build... \033[0m"   # 提示构建镜像
    docker-compose -f ${RUN_ENV_TYPE}_compose.yml build $1   # 构建指定服务镜像
    sleep 3s                                                 # 等待3秒确保构建完成

    # 启动容器
    echo -e "\033[32m Starting container $1 ... \033[0m"     # 提示启动容器
    docker-compose -f ${RUN_ENV_TYPE}_compose.yml up -d $1   # 后台启动指定容器
    sleep 3s                                                 # 等待3秒确保启动完成

    # 查看状态
    echo -e "\033[32m Container run status... \033[0m"       # 提示查看状态
    docker ps -a                                             # 显示所有容器状态
else
    echo -e "\033[32m ${RUN_ENV_TYPE} all... \033[0m"       # 提示更新所有服务

    # 停止所有容器
    echo -e "\033[32m Stopping all containers... \033[0m"    # 提示停止所有容器
    docker-compose -f ${RUN_ENV_TYPE}_compose.yml down       # 停止并移除所有容器
    sleep 3s                                                 # 等待3秒确保操作完成

    # 更新代码
    echo -e "\033[32m Pulling the latest code from the remote... \033[0m"  # 提示拉取代码
    git pull origin main                                    # 从main分支拉取最新代码
    sleep 3s                                                 # 等待3秒确保拉取完成

    # 设置超时
    echo -e "\033[32m Set environment variables(120s)... \033[0m"  # 提示设置超时
    export DOCKER_CLIENT_TIMEOUT=120                         # Docker客户端超时设置
    export COMPOSE_HTTP_TIMEOUT=120                          # Compose HTTP超时设置

    # 构建镜像
    echo -e "\033[32m Images is being build... \033[0m"      # 提示构建镜像
    docker-compose -f ${RUN_ENV_TYPE}_compose.yml build      # 构建所有服务镜像
    sleep 3s                                                 # 等待3秒确保构建完成

    # 启动容器
    echo -e "\033[32m Starting container... \033[0m"         # 提示启动容器
    docker-compose -f ${RUN_ENV_TYPE}_compose.yml up -d      # 后台启动所有容器
    sleep 5s                                                 # 等待5秒确保启动完成

    # 查看状态
    echo -e "\033[32m Container startup status... \033[0m"   # 提示查看启动状态
    docker-compose -f ${RUN_ENV_TYPE}_compose.yml top        # 显示容器进程信息

    echo -e "\033[32m Container run status... \033[0m"       # 提示查看运行状态
    docker ps -a                                             # 显示所有容器状态
fi
