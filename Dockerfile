From golang:1.14rc1-alpine3.11

WORKDIR /go
COPY . .

CMD ["go", "run", "main.go"]
