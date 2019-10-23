FROM golang:latest

#  作者签名
MAINTAINER 977767937@qq.com

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

#时间
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app
COPY . /app
RUN go build -o main .

#  设置对外端口为 9000
EXPOSE 9000

CMD ["/app/main"]

 

