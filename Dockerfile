FROM golang:1.16-buster as builder

# copy source files
COPY ./*.go /sources/
COPY ./go.mod /sources/
COPY ./go.sum /sources/

WORKDIR /sources

# get deps
RUN go get -v -t -d ./...

# build application
RUN go build -v github.com/raver119/geoip_service
RUN mkdir /application
RUN cp /sources/geoip_service /application/service

# Final step. Copy application
FROM ubuntu:20.04

ENV TZ=Europe/London
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt update && apt install -y ca-certificates

RUN mkdir /application

# copy the binary from builder
COPY --from=builder /application/service /application/

# expose default port
EXPOSE 8080

WORKDIR /application
CMD ["./service"]

