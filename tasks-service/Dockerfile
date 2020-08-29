# Golang build image
FROM golang:1.15 as builder

# Maintainer
LABEL maintainer="Pedro Querido <pedroquerido14@gmail.com>"

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/api cmd/api/main.go

####

FROM alpine:latest

COPY --from=builder /app/bin/api .

EXPOSE 8080

ENTRYPOINT ["./api"]