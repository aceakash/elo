FROM golang:1.7

WORKDIR /go/src/github.com/aceakash/elo
ADD . /go/src/github.com/aceakash/elo

RUN go build slack/cmd/main.go
CMD ./main
