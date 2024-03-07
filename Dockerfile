FROM golang:1.21 AS builder

RUN go env -w GOOS=linux
COPY . /usr/local/BlogGoBackEnd
WORKDIR /usr/local/BlogGoBackEnd
RUN go build -o spblog main.go

FROM python:3.11-slim
MAINTAINER Broken_song<Broken_song@outlook.com>
ENV VERSION 0.0.1
RUN apt-get update && apt-get install -y bash supervisor 
WORKDIR /starpearlblog
VOLUME ["/starpearlblog/logs"]
VOLUME ["/starpearlblog/storage"]
COPY --from=builder /usr/local/BlogGoBackEnd/ /starpearlblog
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
ENV LANG C.UTF-8
EXPOSE 8081
CMD ["supervisord", "-n", "-c", "/starpearlblog/supervisord.conf"]

