FROM golang:1.19 as builder

WORKDIR /usr/src/app

COPY go.* ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/local/bin/app ./...

CMD ["app"]