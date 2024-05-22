FROM golang:alpine

WORKDIR "/"

ENV branch $BRANCH
RUN apk add --no-cache curl git bash

RUN curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/install.sh | bash

ENTRYPOINT ["bash", "/root/run_elichika"]
