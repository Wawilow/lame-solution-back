FROM golang:alpine AS builder
WORKDIR /backend

COPY . .
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o ./main ./cmd/server/main.go

FROM alpine:latest
WORKDIR ./src

COPY --from=builder /backend/main ./

EXPOSE $PORT
ENTRYPOINT ["./main" ]