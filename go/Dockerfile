FROM golang:1.17.6-alpine3.15

WORKDIR /go/src/app
ADD ./app .

RUN apk update && \
    apk add --no-cache git && \
    go get github.com/pilu/fresh && \
    go get -u github.com/labstack/echo/v4 && \
    go mod tidy

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8080

CMD ["go", "run", "server.go"]