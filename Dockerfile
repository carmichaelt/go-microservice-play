FROM golang:1.15.6-alpine3.12 AS build-env

RUN apk add --no-cache git

RUN adduser -D -u 10000 carmichaelt
RUN mkfir /microservices/ && chown carmichaelt /microservices/
USER florin


WORKDIR /microservies/
ADD . /microservices/

RUN CGO_ENABLES=0 go build -o /microservices/gcuk .

#final stage
FROM apline:3.8

RUN adduser -D -u 10000 carmichaelt
USER carmichaelt

WORKDIR /
COPY --from=build-env /microservices/certs/docker.localhost.* /
COPY --from=build-env /microservices/gcuk /

EXPOSE 8080

CMD ["/gcuk"]

