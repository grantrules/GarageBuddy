FROM golang:alpine

ENV port=8080

WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN apk add --no-cache make
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux make


CMD ["./output/garagebuddy"]

EXPOSE ${port}
