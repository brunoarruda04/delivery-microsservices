# build the binary
FROM golang:1.22.5-alpine AS builder
ENV APP_BINARY_DIR=build
ENV APP_BINARY=auth-service

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/${APP_BINARY} ./cmd/main.go

# build a tiny docker image
FROM alpine:3.19.0 AS runner
ENV APP_BINARY_DIR=/app/build/auth-service
ENV APP_BINARY=auth-service

RUN apk add --no-cache bash
RUN mkdir /app
WORKDIR /app

COPY --from=builder $APP_BINARY_DIR .

RUN chmod +x $APP_BINARY

CMD /app/$APP_BINARY