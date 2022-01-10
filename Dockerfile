FROM golang:alpine AS build

WORKDIR /usr/src

COPY . .

RUN go get

RUN go build -o ./generator

FROM alpine

COPY --from=build \
    /usr/src/generator \
    /usr/local/bin/

CMD ["/usr/local/bin/generator"]
