# Golang build image
FROM golang:1.15 as build

# Maintainer
LABEL maintainer="Pedro Querido <pedroquerido14@gmail.com>"

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN make install build-api

####

FROM scratch

COPY --from=build /app/bin/api /

EXPOSE 8080

ENTRYPOINT ["/api"]