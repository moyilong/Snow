FROM golang:1.23

ENV GOPROXY=https://ghproxy.cn,direct

COPY . /src

WORKDIR /src

RUN go build .