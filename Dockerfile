FROM golang:1.24.0-alpine AS builder

RUN apk update && apk add ca-certificates git gcc g++ libc-dev binutils tzdata

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /opt

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN /go/bin/swag init -g cmd/main.go

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o bin/application ./cmd/

FROM alpine:3.19 AS runner

RUN apk update && apk add ca-certificates libc6-compat openssh bash tzdata && rm -rf /var/cache/apk/*

ENV TZ=Europe/Moscow

WORKDIR /opt

COPY --from=builder /opt/docs /opt/docs
COPY --from=builder /opt/keys /opt/keys
COPY --from=builder /opt/files /opt/files
COPY --from=builder /opt/pkg /opt/pkg
COPY config.yaml /opt
COPY --from=builder /opt/bin/application ./

CMD ["./application"]
