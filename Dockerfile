FROM golang:alpine as builder

ENV GOPROXY https://goproxy.cn

RUN mkdir /app && mkdir /app/data
WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY cmd/ /app/cmd/
COPY pkg/ /app/pkg/
COPY main.go /app/
RUN go build -o topn


FROM alpine

RUN mkdir /app && mkdir /app/data
WORKDIR /app

COPY config.yaml .
COPY --from=builder /app/topn .
