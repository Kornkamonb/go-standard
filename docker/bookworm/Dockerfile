FROM golang:1.25-bookworm

ARG http_proxy
ARG https_proxy
ARG no_proxy
ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG NO_PROXY

ENV http_proxy=$http_proxy
ENV https_proxy=$https_proxy
ENV no_proxy=$no_proxy
ENV HTTP_PROXY=$HTTP_PROXY
ENV HTTPS_PROXY=$HTTPS_PROXY
ENV NO_PROXY=$NO_PROXY

WORKDIR /app

RUN apt-get update && \
    apt-get install -y libaio1 && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum* ./
RUN go mod download

ENV http_proxy=
ENV https_proxy=
ENV no_proxy=
ENV HTTP_PROXY=
ENV HTTPS_PROXY=
ENV NO_PROXY=
