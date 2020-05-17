FROM golang:1.14-alpine3.11

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go install github.com/gobuffalo/pop/v5/soda

EXPOSE 8080

RUN go get github.com/pilu/fresh
CMD ["fresh"]
