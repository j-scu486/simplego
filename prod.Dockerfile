FROM docker.io/library/golang:1.22.3-bookworm AS builder

RUN mkdir /workdir
COPY . /workdir
WORKDIR /workdir
RUN CGO_ENABLED=0 go build -ldflags='-s -w' -buildvcs=false -o app cmd/main.go

FROM scratch AS production

EXPOSE 800

COPY --from=builder /workdir/app app
