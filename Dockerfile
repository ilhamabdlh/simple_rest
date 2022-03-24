FROM golang:1.17

MAINTAINER ilham

WORKDIR /go/cmd
COPY ./cmd

RUN go get -d -v
RUN go build -v

CMD ["./simple-rest"]