FROM golang:alpine AS builder

RUN apk update && apk add make gcc git libc-dev

COPY . /go/src/gitlab.com/hjames9/simpler
WORKDIR /go/src/gitlab.com/hjames9/simpler
RUN GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
RUN go get -u golang.org/x/lint/golint
RUN make

FROM alpine:latest

COPY --from=builder /go/bin/simpler /go/bin/simpler
ENV PATH="/go/bin:${PATH}"
