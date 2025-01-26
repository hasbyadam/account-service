# Builder
FROM golang:1.20.7-alpine3.17

RUN apk update && apk add --no-cache git

WORKDIR /app

# Download Go modules
COPY . .

# Build
RUN go build -o bin .

COPY .env .

ENTRYPOINT ["./bin"]