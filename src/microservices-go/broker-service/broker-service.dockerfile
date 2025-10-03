#NOTE: base go image

FROM golang:1.24.4-alpine3.20 AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api 

RUN chmod +x /app/brokerApp

#NOTE: build a tiny broker image
FROM alpine:latest

RUN mkdir /app 

COPY --from=builder /app/brokerApp /app 

CMD ["/app/brokerApp"]
