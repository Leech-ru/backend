FROM golang:1.24.0-alpine AS builder

RUN apk update && apk add ca-certificates git gcc g++ libc-dev binutils tzdata

WORKDIR /opt

COPY . .

RUN go mod download && go mod verify

RUN go build -o bin/application ./cmd

FROM alpine:3.19 AS runner

RUN apk update && apk add ca-certificates libc6-compat openssh bash tzdata && rm -rf /var/cache/apk/*

ENV TZ=Europe/Moscow

WORKDIR /opt

COPY --from=builder /opt/docs /opt/docs
COPY --from=builder /opt/keys /opt/keys
COPY config.yaml /opt
COPY --from=builder /opt/bin/application ./

CMD ["./application"]
