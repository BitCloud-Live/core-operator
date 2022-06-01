FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install wget && \
    apt-get clean

WORKDIR /app
RUN wget 'https://coldcdn.com/api/cdn/f2cobx/terminal/v2.5.2/meson-linux-amd64.tar.gz'    # download the terminal package
WORKDIR /app/meson-linux-amd64   
# TODO: run meson network in passive mode?

