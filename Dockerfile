FROM golang:1.20.6-alpine as builder

ENV ROOT=/go/src/
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY ./main.go ${ROOT}
COPY ./crud.go ${ROOT}
COPY ./utils ${ROOT}
COPY go.mod ${ROOT}

RUN go mod tidy
RUN go run .
