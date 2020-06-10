FROM golang:1.10

ADD . /go/src/github.com/tilt-dev/random_number/fe
RUN go install github.com/tilt-dev/random_number/fe
ENTRYPOINT /go/bin/fe
