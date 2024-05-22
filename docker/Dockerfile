FROM golang:alpine

WORKDIR "/"

ARG BRANCH
RUN apk add --no-cache curl git bash

RUN echo $BRANCH
RUN curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/install.sh | bash

ENTRYPOINT ["bash", "/root/run_elichika"]
