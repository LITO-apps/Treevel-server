FROM golang:1.14-alpine3.11

WORKDIR /go
COPY . .

CMD ["go", "run", "main.go"]
