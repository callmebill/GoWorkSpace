FROM golang:latest

MAINTAINER Huang zhigang <http://huangzhigang1024@qq.com>

ENV GOROOT /usr/local/go

ENV GOPATH /GoWorkSpace

ENV GOBIN  $GOPATH/bin

ENV PATH $GOROOT/bin:$PATH

RUN apt-get install -y git

WORKDIR /

RUN git clone git://github.com/callmebill/GoWorkSpace.git

WORKDIR $GOPATH/src/github.com/callmebill/todolist/

RUN go install main.go

ENTRYPOINT $GOPATH/bin/main

EXPOSE 8888