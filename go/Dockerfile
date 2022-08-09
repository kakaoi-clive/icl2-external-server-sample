# build
FROM golang:alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /build
COPY .. .
RUN echo "nobody:x:65534:65534:nobody:/nonexistent:/usr/sbin/nologin" > passwd
RUN go build -ldflags="-s -w" -mod vendor ./cmd/external-server-sample

# deploy
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/passwd /etc/passwd
COPY --from=builder /build/external-server-sample /app/external-server-sample
USER nobody
EXPOSE 8080
ENTRYPOINT ["/app/external-server-sample"]
