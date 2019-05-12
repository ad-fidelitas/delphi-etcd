FROM golang:1.12.5-alpine

WORKDIR /usr/src/app

RUN apk add --no-cache git mercurial \
    && go get github.com/coreos/etcd \
    && apk del git mercurial
RUN echo $GOPATH

#COPY package*.json tsconfig.json swaggerDoc.js /usr/src/app/
#RUN npm install

#COPY src /usr/src/app/src/
#RUN npm run build && rm -rf src

EXPOSE 2379

CMD [ "$GOPATH/bin/etcd" ]

