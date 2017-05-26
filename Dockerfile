FROM golang:1.8.0-alpine

RUN apk update && \
  apk add git glide

WORKDIR /go/src/github.com/nycdavid/dynamornr/

COPY glide* /go/src/github.com/nycdavid/dynamornr/

RUN glide install

ADD ./ ./

ENV PATH "$PATH:/go/bin"

RUN go install
