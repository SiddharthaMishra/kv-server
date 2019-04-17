# Use a multistage build and run process to optimize image size
FROM golang:latest as builder

LABEL maintainer="SiddharthaMishra <sidm1999@gmail.com>"

# Install dep
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

WORKDIR $GOPATH/src/github.com/SiddharthaMishra/kv-server-go
ADD Gopkg.lock ./Gopkg.lock
ADD Gopkg.toml ./Gopkg.toml

RUN dep ensure -vendor-only

ADD *.go ./

# Test and build binary
RUN go test -v && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/kv-server

# Start from fresh

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/bin/kv-server .

CMD [ "./kv-server" ]

EXPOSE 8000