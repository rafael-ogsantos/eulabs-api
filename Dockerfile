FROM golang:1.14.6-alpine3.11

WORKDIR /go/src

ENTRYPOINT ["tail", "-f", "/dev/null"]