FROM golang:1.13.6-alpine3.11
RUN apk update && apk add gcc git
WORKDIR /goelasticsearch/app
COPY ./go.mod /goelasticsearch/go.mod
COPY ./app /goelasticsearch/app
RUN go build ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

ENV ELASTIC_APM_SERVER_URL=http://apm:8200
ENV ELASTIC_APM_SERVICE_NAME=golangs

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /goelasticsearch/app .
CMD ["./app"]
