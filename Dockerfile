FROM golang:1.16 as builder

ADD . /app
WORKDIR /app

RUN go build -o /tmp/monbundle ./cli/main.go



FROM debian:11-slim
COPY --from=builder /tmp/monbundle /monbundle
ENTRYPOINT [ "/monbundle" ]

ENV DISK_DEV sda
