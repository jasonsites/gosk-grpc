FROM golang:1.18

ENV MIGRATE_VERSION=v4.15.1

RUN \
  # just
  curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin \
  && chmod +x /usr/local/bin/just \
  # migrate
  && curl -L https://github.com/golang-migrate/migrate/releases/download/$MIGRATE_VERSION/migrate.linux-amd64.tar.gz | tar xvz \
  && mv migrate /usr/bin/migrate \
  && chmod +x /usr/bin/migrate \
  # air
  && curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
  && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air \
  # gotestsum
  && curl -sSL https://github.com/gotestyourself/gotestsum/releases/download/v1.6.4/gotestsum_1.6.4_linux_amd64.tar.gz | tar -xz -C /usr/local/bin gotestsum

ENV APP=/app

RUN mkdir -p $APP
WORKDIR $APP
# COPY go.mod go.sum $APP/
COPY . $APP
RUN go mod vendor && go mod tidy
# RUN just lint

CMD just server
