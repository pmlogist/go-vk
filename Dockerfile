FROM golang:1.17.3 as builder

ENV GOBIN=$GOPATH
ENV PATH=$PATH:$GOBIN
