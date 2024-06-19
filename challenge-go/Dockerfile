FROM golang:1.22 as builder

RUN apt-get update && apt-get install -y musl-dev musl-tools gcc
RUN mkdir /challenge-go
WORKDIR /challenge-go
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . .
RUN swag init -g cmd/main.go --output /challenge-go/docs
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .bin/main cmd/main.go


FROM alpine:latest as cacerts
RUN apk --no-cache add ca-certificates


FROM alpine:latest
ARG APP_ENV
ENV APP_ENV=$APP_ENV

WORKDIR /challenge-go
COPY --from=builder /challenge-go /challenge-go
COPY --from=cacerts /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip

ENV APP_ENV=$APP_ENV
ENV TZ America/Lima

EXPOSE 8080

CMD ["./.bin/main" ]



