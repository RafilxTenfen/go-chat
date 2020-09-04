FROM golang:1.14 as stage

WORKDIR /app
COPY . .
RUN make GOOS=linux GOARCH=amd64 CGO_ENABLED=0 EXTLDFLAGS="-static" FLAGS="-tags netgo" bot

FROM alpine as gorobot

WORKDIR /app

RUN apk --no-cache add ca-certificates

ENV QUANTITY_MESSAGE_QUEUE 50
ENV RABBIT_MQ_URL amqp://guest:guest@localhost:5672/
ENV DATABASE_HOST localhost
ENV DATABASE_PORT 5432     
ENV DATABASE_USER postgres 
ENV DATABASE_PASSWORD postgres 
ENV DATABASE_NAME dbgochat

COPY --from=stage /app/bot /app
CMD ["./bot", "run", "-q", "queue-name,q1", "--useshell=false"]