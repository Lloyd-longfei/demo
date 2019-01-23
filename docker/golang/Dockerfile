FROM hub.c.163.com/library/golang:latest

MAINTAINER anonymous "test@163.com"

RUN echo "Asia/Shanghai" > /etc/timezone
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN dpkg-reconfigure -f noninteractive tzdata



WORKDIR $GOPATH/src/demo




