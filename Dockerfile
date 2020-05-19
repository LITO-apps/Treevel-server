FROM golang:1.14-alpine3.11

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go install github.com/gobuffalo/pop/v5/soda

EXPOSE 8080

RUN go get github.com/pilu/fresh
CMD ["fresh"]
