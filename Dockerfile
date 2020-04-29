FROM golang:1.14-alpine3.11

WORKDIR /go
COPY . .

RUN go build -mod=readonly -o app main.go

CMD ["./app"]
