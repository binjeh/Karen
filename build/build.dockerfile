#
# Karen - A highly efficient, multipurpose Discord bot written in Golang
#
# Copyright (C) 2015-2017 Lukas Breuer
# Copyright (C) 2017 Subliminal Apps
#
# This file is a part of the Karen Discord-Bot Project ("Karen").
#
# Karen is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License,
# or (at your option) any later version.
#
# Karen is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
#
# See the GNU Affero General Public License for more details.
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <http://www.gnu.org/licenses/>.
#

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
