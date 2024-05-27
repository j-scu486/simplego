FROM docker.io/library/golang:1.22.3-bookworm

EXPOSE 8000

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
