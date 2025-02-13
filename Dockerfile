ARG VERSION_GOLANG="1.23"
ARG VERSION_ALPINE="3.21"

FROM golang:${VERSION_GOLANG}-alpine${VERSION_ALPINE} AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . .

RUN go build -o main ./cmd/

FROM alpine:${VERSION_ALPINE}

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
