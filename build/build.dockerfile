FROM golang:alpine

RUN apk add --no-cache ca-certificates curl
RUN update-ca-certificates

RUN apk add --no-cache --virtual .build-deps git gnupg alpine-sdk bash

RUN curl https://glide.sh/get | sh

WORKDIR /build-deps
RUN \
    curl -L https://github.com/logological/gpp/releases/download/2.25/gpp-2.25.tar.bz2 -o gpp-2.25.tar.bz2 && \
    tar xvfj gpp-2.25.tar.bz2 && \
    cd gpp-2.25 && \
    ./configure && \
    make && make install

COPY bootstrap.sh /build-deps
RUN ./bootstrap.sh

WORKDIR /
