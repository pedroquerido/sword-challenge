FROM golang:1.15 as builder

LABEL maintainer="Pedro Querido <pedroquerido14@gmail.com>"

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/auth cmd/main.go

####

FROM alpine:latest

COPY --from=builder /app/bin/auth .

EXPOSE 8080

ENTRYPOINT ["./auth"]