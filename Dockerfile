# BUILD
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./*.go ./

RUN go build -o /bug-tracker

##
## Deploy
##
FROM alpine:3.10

WORKDIR /

COPY --from=build /bug-tracker /bug-tracker

EXPOSE 8080

ENTRYPOINT ["/bug-tracker"]
