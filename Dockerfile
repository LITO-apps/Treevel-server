FROM golang:1.14-alpine3.11

WORKDIR /go/src/app

COPY . .

RUN go build -mod=readonly -o app main.go

EXPOSE 8080

CMD ["./app"]
