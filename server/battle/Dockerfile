FROM golang:1.7.4-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash git openssh && \
    go get github.com/FiloSottile/gvt && \
    mkdir -p /go/src/github.com/shiwano/submarine/server/battle
WORKDIR /go/src/github.com/shiwano/submarine/server/battle

COPY vendor /go/src/github.com/shiwano/submarine/server/battle/vendor
RUN gvt restore

COPY config.json /go/src/github.com/shiwano/submarine/server/battle
COPY . /go/src/github.com/shiwano/submarine/server/battle
RUN go build

ARG env=development
ENV SUBMARINE_ENV $env

EXPOSE 5000
ENTRYPOINT ["./battle"]
