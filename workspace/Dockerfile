FROM golang:1.11-stretch

VOLUME /usr/src/logger

WORKDIR /root

RUN apt-get update \
	&& apt-get install -y \
		vim \
	&& apt-get clean all

ENV TERM xterm
ENV GOPATH /go:/usr
ENV GO111MODULE on

COPY ./.bashrc /root/.bashrc

WORKDIR /usr/src/logger

# CMD ["/run.sh"]