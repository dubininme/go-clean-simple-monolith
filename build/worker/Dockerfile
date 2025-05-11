FROM golang:1.23-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/worker ./cmd/worker

FROM alpine:3.19

COPY --from=builder /go/bin/worker /worker

ENTRYPOINT ["/worker"]
