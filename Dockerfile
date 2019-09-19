FROM golang:1.13 as builder

WORKDIR /src
COPY . /src

RUN go mod vendor && \
    CGO_ENABLED=0 go build -v -o notifier main.go

FROM alpine:3.10

COPY LICENSE README.md /
COPY --from=builder /src/notifier /notifier

ENTRYPOINT ["/notifier"]
