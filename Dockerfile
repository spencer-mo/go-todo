FROM golang:1.17-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    git \
    && update-ca-certificates


FROM base AS dev
WORKDIR /app

# install Air for hot reloading
RUN go get -u github.com/cosmtrek/air && go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 5001
EXPOSE 2345

# run air (hot reloading)
ENTRYPOINT ["air"]

FROM base AS builder
WORKDIR /app

COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o todo -a .

FROM alpine:latest as prod

COPY --from=builder /app/todo /usr/local/bin/todo
EXPOSE 5001

ENTRYPOINT ["/usr/local/bin/todo"]
