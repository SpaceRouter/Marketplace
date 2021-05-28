FROM golang

RUN apt update && apt install libpam0g-dev -y

COPY src /source
WORKDIR /source

RUN go get
RUN go build -o /usr/bin/marketplace

RUN mkdir /config && cp config/*.yaml /config -r

WORKDIR /

ENV GIN_MODE=release

CMD marketplace Server