FROM golang:alpine as build 
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . ./
RUN go build -o /main src/cmd/*.go

FROM alpine:latest
WORKDIR /
COPY /migrations /migrations
COPY --from=build /main /

ENTRYPOINT [ "/main" ]