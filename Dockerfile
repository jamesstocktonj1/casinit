# syntax=docker/dockerfile:1

# build container
FROM golang:1.22 AS builder

WORKDIR /app

COPY . ./
RUN go mod download

ARG CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o ./casinit cmd/casinit/main.go


# run container
FROM scratch

WORKDIR /app

COPY --from=builder /app/casinit ./casinit

ENTRYPOINT ["./casinit"]