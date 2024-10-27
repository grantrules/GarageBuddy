FROM golang:alpine AS build-stage

ENV port=8080

WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN apk add --no-cache make
RUN go mod download

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux make

FROM alpine

COPY --from=build-stage /app/output /app
COPY --from=build-stage /go/bin/dlv /

EXPOSE ${port}

CMD "/app/server"