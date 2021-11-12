FROM golang:alpine

ENV GIN_MODE=release
ENV APP_PORT=3004
ENV GIN_PROXY_PORT=80

WORKDIR /go/src/docker-sync

COPY . /go/src/docker-sync

RUN go get github.com/codegangsta/gin

RUN go build

EXPOSE $APP_PORT
EXPOSE $GIN_PROXY_PORT

CMD ["gin","-i","-a","3004","-p","80","run"]