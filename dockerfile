FROM golang:1.15.2-alpine3.12

RUN mkdir /app

ADD . /app

WORKDIR /app
EXPOSE 8080

RUN go build -o main .

CMD ["/app/main"]