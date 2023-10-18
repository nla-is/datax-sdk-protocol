FROM golang:1.21-bookworm

RUN mkdir /src
WORKDIR /src

ADD stub.go .

RUN go build -buildmode=c-shared -o /libdatax-sdk.so stub.go

ENTRYPOINT ["cp", "/libdatax-sdk.h", "/libdatax-sdk.so", "/mnt"]