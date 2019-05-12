FROM golang:1.12.5

WORKDIR /usr/src/app

# RUN apk add --no-cache git mercurial \
RUN go get github.com/coreos/etcd 
# && apk del git mercurial
# ARG GOPATH=$GOPATH

ENV GOPATH $GOPATH

EXPOSE 2379


CMD ["sh", "-c", "${GOPATH}/bin/etcd"]