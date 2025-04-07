FROM golang:1.23.5-alpine3.21 AS builder

COPY . /app
WORKDIR /app

RUN go get
RUN go build -o /app/api

FROM alpine:3.21

COPY --from=builder /app/api /app/api

EXPOSE 3000

CMD ["/app/api"]
