FROM golang:1.23 as build

ENV GOPROXY=https://ghproxy.cn,direct

COPY . /src

WORKDIR /src

RUN go build ./cmd/gossip

FROM scratch

COPY --from=build --chmod=0777 /src/gossip /gossip

ENTRYPOINT [ "/gossip" ]
