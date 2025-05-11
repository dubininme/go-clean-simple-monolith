FROM golang:1.23-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./cmd/api

FROM alpine:3.19

COPY --from=builder /go/bin/app /app

EXPOSE 80

ENTRYPOINT ["/app"]
