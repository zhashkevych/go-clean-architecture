FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY ./.bin/app .
COPY ./config/ ./config/