# 多阶段构建
#构建一个 builder 镜像，目的是在其中编译出可执行文件mck
#构建时需要将此文件放到代码根目录下
FROM golang:alpine  as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY https://goproxy.cn,direct
#将上层整个文件夹拷贝到/build
ADD . /build/src
WORKDIR /build/src
#COPY --from=builder  /build/src/config/config.yml /app/config/config.yml
#去掉了调试信息 -ldflags="-s -w" 以减小镜像尺寸
RUN   go build -ldflags="-s -w"  -o upserver ./cmd/upserver/main.go

FROM alpine
#更新软件源
RUN apk update --no-cache && apk add --no-cache tzdata \
 && apk add --no-cache docker-cli
#设置本地时区，这样我们在日志里看到的是北京时间了
ENV TZ Asia/Shanghai
#安装 docker
#RUN apk add --no-cache docker-cli
#安装helm
#COPY ./helm /usr/local/bin/helm
WORKDIR /home
#安装push插件
COPY ./helm-push_0.10.3_linux_amd64.tar.gz ./
RUN mkdir -p /root/.local/share/helm/plugins/helm-push \
 && tar zxvf  helm-push_0.10.3_linux_amd64.tar.gz  -C /root/.local/share/helm/plugins/helm-push \
 && rm ./helm-push_0.10.3_linux_amd64.tar.gz

COPY --from=builder  /build/src/config/config.yml /home/config/config.yml
COPY --from=builder /build/src/upserver /home/upserver
#COPY /var/run/docker.sock /var/run/docker.sock

CMD ["./upserver"]
EXPOSE 9680

#需要设置映射 /var/run/docker.sock:/var/run/docker.sock 将使用宿主机中的docker
#映射 /usr/local/bin/helm:/usr/local/bin/helm 使用helm
