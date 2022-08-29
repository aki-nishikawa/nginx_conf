FROM golang:1.18.4-alpine3.16

COPY ./src /go/src

WORKDIR /go/src/chat

RUN go build

CMD ["./chat"]
