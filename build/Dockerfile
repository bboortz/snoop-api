#
# stage 1
# 
FROM golang:latest as build

WORKDIR /
RUN mkdir -p /go/src/github.com/bboortz/snoop-api
WORKDIR /go/src/github.com/bboortz/snoop-api
COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

RUN go get github.com/ghodss/yaml && \
  CGO_ENABLED=0 GOOS=linux go build -a -tags "netgo static_build" -installsuffix netgo -ldflags "-w -s" -o snpop cmd/main.go



#
# stage 2
#
FROM scratch
LABEL maintainer "Benjamin Boortz <benjamin.boortz@secure.mailbox.org>"

WORKDIR /
COPY examples /examples
COPY examples/snoop.yaml /
COPY --from=build go/src/github.com/bboortz/snoop-api/snoop /
ENTRYPOINT ["/snoop"]
