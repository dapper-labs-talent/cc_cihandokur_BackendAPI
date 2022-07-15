FROM golang:1.18-alpine3.16 AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app/
COPY . /app/
RUN go build -o /app/devlab ./cmd


FROM alpine
RUN mkdir -p app
WORKDIR /app/
COPY --from=builder /app/devlab/ /app/
COPY --from=builder /app/config/config.toml /app/config/config.toml

COPY docker-entrypoint.sh /app
RUN chmod +x docker-entrypoint.sh

EXPOSE 9090

ENTRYPOINT ["./docker-entrypoint.sh"]

