FROM golang:latest
 
MAINTAINER Huang zhigang <http://huangzhigang1024@qq.com>

ENV GOROOT /usr/local/go

ENV GOPATH /goworkspace

ENV PATH $GOROOT/bin:$PATH

ADD /var/jenkins_home/workspace/todo &GOPATH

ENTRYPOINT $GOPATH/bin/todos

EXPOSE 8888
