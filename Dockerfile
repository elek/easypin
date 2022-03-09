FROM alpine
WORKDIR /opt/easypin
ADD cmd/easypin/easypin /usr/bin/easypin
ADD web/dist web/dist
ENTRYPOINT ["/usr/bin/easypin"]
