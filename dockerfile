FROM golang:1.18-bullseye

RUN apt-get update \
  && apt-get install -y --no-install-recommends chromium \
  && apt-get clean

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

CMD [ "./go-web-scraping" ]
