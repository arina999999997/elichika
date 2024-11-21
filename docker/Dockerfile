FROM docker.io/library/debian:latest AS builder

# First - build
ARG GO_VERSION=1.22.5

ARG TARGETARCH TARGETOS

ENV CGO_ENABLED=1

RUN apt update && apt install -y curl git gcc
RUN curl -LO https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz

WORKDIR /elichika/

COPY ./ ./

# Replace database storage paths
RUN sed -i 's|userdata.db|/data/userdata.db|g' config/config.go
RUN sed -i 's|serverdata.db|/data/serverdata.db|g' config/config.go

RUN sed -i 's|./config.json|/data/config.json|g' config/config.go
RUN sed -i 's|./config.json|/data/config.json|g' config/runtime.go

RUN export PATH=$PATH:/usr/local/go/bin && \
    GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build

# Not currently possible - elichika uses too much ram
#RUN rm -rf assets/ && git clone https://github.com/arina999999997/harasho assets --depth 1

#RUN ./elichika init

## Second - sort stuff idk

FROM --platform=$TARGETPLATFORM docker.io/library/debian:bookworm-slim

ARG TARGETPLATFORM

RUN mkdir -p /root/elichika/
COPY --from=builder /elichika/elichika /root/elichika/elichika

RUN mkdir -p /data/

RUN apt update && apt install git -y

WORKDIR /root/elichika/

RUN git clone https://github.com/arina999999997/harasho assets --depth 1

#COPY --from=builder /elichika/assets/ /root/elichika/assets/

COPY --from=builder /elichika/webui/ /root/elichika/webui/
RUN rm -rf webui/*.go && rm -rf webui/*/*.go

COPY --from=builder /elichika/publickey.pem /root/elichika/publickey.pem
COPY --from=builder /elichika/privatekey.pem /root/elichika/privatekey.pem

ARG folder="server init jsons"
COPY --from=builder /elichika/${folder}/ /root/elichika/${folder}/

ENTRYPOINT ["/root/elichika/elichika", "rebuild_assets", "keep_alive"]
