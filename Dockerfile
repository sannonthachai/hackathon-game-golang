FROM golang:1.16.6-alpine3.14 AS build-stage

ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /go/src/app/
COPY . .

RUN go mod download
RUN go build -a -installsuffix cgo -o find-the-hidden-backend .

FROM alpine:latest AS production-stage

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add --update --no-cache socat curl tzdata findutils
RUN ln -fs /usr/share/zoneinfo/Asia/Bangkok /etc/localtime

WORKDIR /root/

COPY --from=build-stage /go/src/app/find-the-hidden-backend /bin/find-the-hidden-backend

ENTRYPOINT [ "find-the-hidden-backend" ]
