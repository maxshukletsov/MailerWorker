FROM golang:alpine as builder

RUN set -ex &&\
    apk add --no-progress --no-cache \
    gcc \
    musl-dev

WORKDIR /src

COPY go.mod /src
COPY go.sum /src

RUN go mod download

COPY . /src

RUN GO111MODULE=on GOOS=linux go build -a  -tags musl

FROM alpine

COPY --from=builder /src/Infrastructure.MailerWorker /
COPY --from=builder /src/.env /

CMD ["/Infrastructure.MailerWorker"]
