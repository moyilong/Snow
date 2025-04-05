FROM golang:1.23 as build

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

COPY . /src

WORKDIR /src

RUN go build -v -o /snow ./cmd

FROM ubuntu:latest

WORKDIR /

COPY --from=build --chmod=0777 /snow /snow
COPY ./config/config.yml /config.yml

ENV TZ=Asia/Shanghai

ENTRYPOINT [ "/snow" ]
