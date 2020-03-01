FROM golang:1.13.6-alpine3.11

COPY ./go.mod /goelasticsearch/go.mod
COPY ./app /goelasticsearch/app

RUN apk add --update --no-cache gcc git bind-dev && rm -rf /var/cache/apk/*
RUN go get -u github.com/oxequa/realize

ENV ELASTIC_APM_SERVER_URL=http://apm:8200
ENV ELASTIC_APM_SERVICE_NAME=golangs

WORKDIR /goelasticsearch/app

RUN CGO_ENABLED=0 go mod tidy
