FROM golang:1.25.0 AS builder

#RUN apt-get update && apt-get install build-essential clang -y

WORKDIR /go/bin/

COPY . .

RUN go build -o /go/bin/grpccreator *.go


FROM ubuntu

RUN apt-get update && apt-get install nano iputils-ping -y

COPY --from=builder /go/bin/grpccreator /go/bin/grpccreator
COPY --from=builder /go/bin/data/ /data

WORKDIR /go/bin
