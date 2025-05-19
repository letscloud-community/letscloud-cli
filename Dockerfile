FROM golang:1.24-alpine AS build

ARG GOARCH=amd64
ENV OUT_D /out

RUN mkdir -p /out
RUN  apk add --update  --no-cache \
     bash \
     coreutils \
     git \
     libc6-compat \
     make

RUN mkdir -p /go/src/github.com/letscloud-community/letscloud-cli
ADD . /go/src/github.com/letscloud-community/letscloud-cli

RUN cd /go/src/github.com/letscloud-community/letscloud-cli && \
    make build GOARCH=$GOARCH

FROM alpine:latest

RUN apk add --update --no-cache \
        ca-certificates \
        libc6-compat \
        openssh

COPY --from=build /go/src/github.com/letscloud-community/letscloud-cli/letscloud /usr/local/bin/letscloud

RUN adduser -D user
USER user:user

ENTRYPOINT ["letscloud"]