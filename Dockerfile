FROM golang:1.26.0-alpine3.23 AS dev

RUN apk update &&\
    apk add git