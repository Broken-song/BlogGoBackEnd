FROM golang:1.21
MAINTAINER Broken_song<Broken_song@outlook.com>

RUN go env -w GOOS=linux
RUN cd /usr/local
RUN git clone https://github.com/Broken-song/BlogGoBackEnd.git
RUN cd BlogGoBackEnd
RUN go build -o starpearl-blog main.go

EXPOSE 8081

CMD ["nohup ./starpearl-blog > ./logs/console.log 2>&1 &"]