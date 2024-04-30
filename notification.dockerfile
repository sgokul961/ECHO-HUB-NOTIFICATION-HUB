FROM golang:1.21-alpine3.19 AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg

RUN go build -o notification ./cmd

FROM alpine:3.19

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/notification .
COPY --from=builder /app/pkg/config/envs/dev.env pkg/config/envs/dev.env

CMD ["sh","-c","echo $PORT && ./notification"]