FROM golang:1.18-alpine AS build

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN mkdir data
RUN mkdir find
RUN mkdir javascript
RUN mkdir scraping
RUN mkdir static
RUN mkdir web

COPY data/*.go data
COPY find/*.go find
COPY javascript/*.go javascript
COPY scraping/*.go scraping
COPY static/*.go static
COPY web/*.go web
COPY main.go .

RUN go build -o go-web-scraping main.go

FROM alpine:3.16

RUN apk update \
    && apk upgrade \
    && apk add --no-cache chromium

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/go-web-scraping go-web-scraping

CMD [ "./go-web-scraping" ]
