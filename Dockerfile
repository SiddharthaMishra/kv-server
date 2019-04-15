FROM golang:latest

LABEL maintainer="SiddharthaMishra <sidm1999@gmail.com>"

WORKDIR $GOPATH/src/github.com/SiddharthaMishra/kv-server-go

RUN sudo apt install go-dep

ADD Gopkg.lock ./Gopkg.lock
ADD Gopkg.toml ./Gopkg.toml

RUN dep ensure -vendor-only

ADD *.go ./

RUN go build -o kvserver .

CMD [ "./kvserver" ]