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

FROM sn0w/karen-build

VOLUME /karen
EXPOSE 1337

WORKDIR /build-deps
RUN cd gpp-2.25 && make uninstall

RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl.sig -o /tmp/youtube-dl.sig
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/bin/youtube-dl
RUN gpg --keyserver eu.pool.sks-keyservers.net --recv-keys DB4B54CBA4826A18
RUN gpg --keyserver eu.pool.sks-keyservers.net --recv-keys 2C393E0F18A9236D
RUN gpg --verify /tmp/youtube-dl.sig /usr/bin/youtube-dl
RUN chmod a+rx /usr/bin/youtube-dl
RUN rm /tmp/youtube-dl.sig

RUN apt-get remove -y git gnupg build-essential
RUN apt-get install -y python ffmpeg

WORKDIR /karen
ENTRYPOINT /karen/karen
