FROM --platform=linux/amd64 golang:1.19-alpine AS builder

RUN set -ex &&\
    apk add --no-progress --no-cache \
      gcc \
      musl-dev


WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .
RUN go get -d -v
RUN go build -a -v -tags musl -o /triage

CMD [ "/triage" ]