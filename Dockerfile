FROM golang:1.20-alpine3.17

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
