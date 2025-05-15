ARG APP_PLATFORM=amd64
ARG GITHUB_USER
ARG GITHUB_TOKEN

FROM golang:1.23-alpine AS dev

ARG APP_PLATFORM
ARG GITHUB_USER
ARG GITHUB_TOKEN

RUN [ "$APP_PLATFORM" = "amd64" ] || [ "$APP_PLATFORM" = "arm64" ] || (echo "Error: APP_PLATFORM should be either amd64 or arm64" >&2; exit 1)
RUN [ -n "$GITHUB_USER" ] || (echo "Error: GITHUB_USER must be specified" >&2; exit 1)
RUN [ -n "$GITHUB_TOKEN" ] || (echo "Error: GITHUB_TOKEN must be specified" >&2; exit 1)

WORKDIR /go/src/app
COPY . /go/src/app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=${APP_PLATFORM}
ENV PATH="${PATH}:/go/bin/linux_${APP_PLATFORM}"

RUN apk add --no-cache git build-base curl && \
    echo "machine github.com login ${GITHUB_USER} password ${GITHUB_TOKEN}" > ~/.netrc && \
    go mod download -x && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /go/bin v1.56.2 && \
    go install github.com/google/wire/cmd/wire@latest && \
    go install github.com/cespare/reflex@latest && \
    go install github.com/golang/mock/mockgen@v1.6.0 && \
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 && \
    go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 80 2345

# CMD задаётся через docker-compose для api или worker 