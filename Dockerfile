FROM golang

ADD . /go/src/github.com/bsnchan/sample-app

RUN go install github.com/bsnchan/sample-app

ENTRYPOINT /go/bin/sample-app

EXPOSE 8080
