FROM golang:1.8.3-stretch

RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y ca-certificates curl git gnupg bash build-essential

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
