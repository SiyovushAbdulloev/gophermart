FROM golang:1.22.12-alpine

RUN apk add --no-cache git curl

RUN go install github.com/cosmtrek/air@v1.43.0

WORKDIR /app
CMD ["air"]
