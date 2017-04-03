FROM golang:alpine

VOLUME /go/src/code.lukas.moe/x/karen
EXPOSE 1337

# Install basic utils
RUN apk add --no-cache --virtual .build-deps curl git wget gnupg

# Install depdencies for running karen
RUN apk add --no-cache --virtual .karen-deps python py-setuptools ffmpeg

RUN wget https://yt-dl.org/downloads/latest/youtube-dl.sig -O /tmp/youtube-dl.sig
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/bin/youtube-dl
RUN gpg --keyserver pool.sks-keyservers.net --recv-keys DB4B54CBA4826A18
RUN gpg --keyserver pool.sks-keyservers.net --recv-keys 2C393E0F18A9236D
RUN gpg --verify /tmp/youtube-dl.sig /usr/bin/youtube-dl
RUN chmod a+rx /usr/bin/youtube-dl
RUN rm /tmp/youtube-dl.sig

# Install dependencies for building karen
RUN apk add --no-cache --virtual .build-deps alpine-sdk bash

RUN curl https://glide.sh/get | sh
RUN go get -v golang.org/x/tools/cmd/goimports
RUN	go get -v github.com/lestrrat/go-bindata/...
RUN go get -v git.lukas.moe/sn0w/ropus

WORKDIR /build
RUN \
    wget https://github.com/logological/gpp/releases/download/2.25/gpp-2.25.tar.bz2 && \
    tar xvfj gpp-2.25.tar.bz2 && \
    cd gpp-2.25 && \
    ./configure && \
    make && make install

WORKDIR /go/src/code.lukas.moe/x/karen
COPY . /go/src/code.lukas.moe/x/karen
RUN ./configure
RUN glide install
RUN make compile

WORKDIR /build
RUN cd gpp-2.25 && make uninstall
RUN apk del .build-deps

WORKDIR /go/src/code.lukas.moe/x/karen
ENTRYPOINT /go/src/code.lukas.moe/x/karen/karen
