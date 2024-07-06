FROM docker.io/library/debian:latest AS builder

# First - build

RUN apt update && apt install -y curl git
RUN curl -LO https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz

WORKDIR /elichika/

COPY ./ ./

# Replace database storage paths
RUN sed -i 's|userdata.db|/data/userdata.db|g' config/config.go
RUN sed -i 's|serverdata.db|/data/serverdata.db|g' config/config.go

RUN sed -i 's|./config.json|/data/config.json|g' config/config.go
RUN sed -i 's|./config.json|/data/config.json|g' config/runtime.go

RUN export PATH=$PATH:/usr/local/go/bin && go build

# Not currently possible - elichika uses too much ram
#RUN rm -rf assets/ && git clone https://github.com/arina999999997/harasho assets --depth 1

#RUN ./elichika init

## Second - sort stuff idk

FROM docker.io/library/debian:bookworm-slim

RUN mkdir -p /root/elichika/
COPY --from=builder /elichika/elichika /root/elichika/elichika

RUN mkdir -p /data/

RUN apt update && apt install git -y

WORKDIR /root/elichika/

RUN git clone https://github.com/arina999999997/harasho assets --depth 1

#COPY --from=builder /elichika/assets/ /root/elichika/assets/

COPY --from=builder /elichika/webui/ /root/elichika/webui/
RUN rm -rf webui/*.go && rm -rf webui/*/*.go

ARG folder="server init jsons"
COPY --from=builder /elichika/${folder}/ /root/elichika/${folder}/

ENTRYPOINT ["/root/elichika/elichika"]
