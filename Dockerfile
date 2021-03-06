FROM golang:1.6-alpine
RUN apk add --update git && rm -rf /var/cache/apk/*
RUN mkdir -p /go/src/github.com/tanji/replication-manager
WORKDIR /go/src/github.com/tanji/replication-manager
COPY . /go/src/github.com/tanji/replication-manager/
RUN go install .
RUN mkdir -p /etc/replication-manager && mkdir -p /usr/share/replication-manager/dashboard
COPY etc/config.toml.sample /etc/replication-manager/
COPY dashboard/* /usr/share/replication-manager/dashboard/
RUN rm -rf /go/src
WORKDIR /go/bin
ENTRYPOINT ["replication-manager"]
CMD ["monitor", "--daemon", "--http-server"]
EXPOSE 10001
