FROM golang:alpine

ENV APP_NAME marketplace

RUN apk add gcc musl-dev && \
    go get -u github.com/swaggo/swag/cmd/swag

COPY src /source
WORKDIR /source

RUN go get && \
 swag init && \
 go build -o /usr/bin/$APP_NAME && \
 rm -rf $GOPATH/pkg/


RUN mkdir /config && cp config/*.yaml /config -r

WORKDIR /

ENV GIN_MODE=release

CMD $APP_NAME