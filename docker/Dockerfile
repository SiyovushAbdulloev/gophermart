FROM golang:1.22.12-alpine3.21 as builder

WORKDIR /gophermart

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /gophermart/gophermart ./cmd/gophermart

FROM scratch
COPY --from=builder /gophermart/gophermart /gophermart

EXPOSE 8080
ENTRYPOINT ["/gophermart"]