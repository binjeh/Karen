FROM sn0w/karen-build

VOLUME /karen
EXPOSE 1337

WORKDIR /build-deps
RUN cd gpp-2.25 && make uninstall

RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl.sig -o /tmp/youtube-dl.sig
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/bin/youtube-dl
RUN gpg --keyserver pool.sks-keyservers.net --recv-keys DB4B54CBA4826A18
RUN gpg --keyserver pool.sks-keyservers.net --recv-keys 2C393E0F18A9236D
RUN gpg --verify /tmp/youtube-dl.sig /usr/bin/youtube-dl
RUN chmod a+rx /usr/bin/youtube-dl
RUN rm /tmp/youtube-dl.sig

RUN apk del .build-deps
RUN apk add --no-cache --virtual .karen-deps python py-setuptools ffmpeg

WORKDIR /karen
ENTRYPOINT /karen/karen
