FROM golang:1.18-alpine AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 \
    go build -mod=readonly -o out/bin/client ./cmd/client && \
    go build -mod=readonly -o out/bin/gateway ./cmd/gateway && \
    go build -mod=readonly -o out/bin/server ./cmd/server


FROM alpine

RUN mkdir -p app
WORKDIR /app
COPY --from=build /src/package.json /src/config/config.toml /src/out/bin/server /src/out/bin/gateway /app/
COPY --from=build /src/internal/migrations /app/internal/migrations
RUN chmod +x config.toml package.json && \
    apk --no-cache add curl supervisor

COPY supervisord.conf /etc/supervisord.conf

ENV MIGRATE_VERSION=v4.15.1
RUN \
  # migrate
  curl -L https://github.com/golang-migrate/migrate/releases/download/$MIGRATE_VERSION/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/bin/migrate && \
    chmod +x /usr/bin/migrate && \
    # enable gRPC health probe
    GRPC_HEALTH_PROBE_VERSION=v0.4.5 && \
    curl -L https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 --output /usr/local/bin/grpc_health_probe && \
    chmod +x /usr/local/bin/grpc_health_probe

EXPOSE 50051
EXPOSE 9051

CMD ["/usr/bin/supervisord"]
