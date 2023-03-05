FROM golang:1.19.3-alpine AS builder
WORKDIR /app

RUN apk add gcc g++

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY internal/ ./internal/
COPY main.go ./main.go

RUN GOOS=linux CGO_ENABLED=1 go build -ldflags="-extldflags=-static" -o import ./main.go

FROM scratch
WORKDIR /app

COPY --from=builder /app /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/import"]
CMD [ "--help" ]
