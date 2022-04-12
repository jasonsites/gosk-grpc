# gosk-grpc
Go Starter Kit for gRPC API Applications

## Architecture

<p align="center">
  <img src="./architecture.png" width="70%" alt="architecture" />
</p>

## Installation
Clone the repository
```sh
$ git clone git@github.com:jasonsites/gosk-grpc.git
$ cd gosk-grpc
```

## Development
**Prerequisites**

- *[Docker Desktop](https://www.docker.com/products/docker-desktop)*
- *[Go 1.18+](https://golang.org/doc/install)*
- *[Protocol Buffers](https://github.com/protocolbuffers/protobuf/)*
- *[Go bindings for Protocol Buffers](https://github.com/golang/protobuf)*

<!-- TODO: Go environment setup documentation

# install and configure go
$ brew install golang (mac)
$ echo 'export GOPATH="$HOME/go"' >> ~/.zshrc
$ echo 'export GOROOT="$(brew --prefix golang)/libexec"' >> ~/.zshrc
$ echo 'export PATH="$GOPATH/bin:$GOROOT/bin:$PATH"' >> ~/.zshrc

# install protobuf compiler
$ brew install protobuf (mac)

# install protoc-gen-go plugin
$ go get google.golang.org/protobuf/cmd/protoc-gen-go

# install protoc-gen-go-grpc plugin
$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# for http/1 gateway, install protoc-gen-grpc-gateway
$ go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

-->

**Show all commands**
```sh
$ docker compose run --rm api just
```

### Migrations
**Run all up migrations**
```sh
$ docker compose run --rm api just migrate
```

**Run up migrations {n} steps**
```sh
$ docker compose run --rm api just migrate-up svcdb {n}
```

**Run down migrations {n} steps**
```sh
$ docker compose run --rm api just migrate-down svcdb {n}
```

### Server
**Run the gRPC server in development mode**
```sh
$ docker compose run --rm --service-ports api
```

**Run the go client for sending basic RPCs**
```sh
$ go run cmd/client/client.go
```

### Testing
**Run migrations on test database**
```sh
$ docker compose run --rm api just migrate-up testdb
```

**Run full test suite**
```sh
$ docker compose run -e POSTGRES_DB=testdb api just test
```

**Run full test suite with code coverage**
```sh
$ docker compose run -e POSTGRES_DB=testdb api just coverage
```

## Building
**Compile server binary**
```sh
$ go build -mod vendor -o out/bin/domain ./cmd/server
```

**Generate go protocol buffers files**
```sh
$ protoc -I=protos --go_out=internal --go-grpc_out=internal protos/*.proto
```

**Generate go protocol buffers files with HTTP/1 gateway**
```sh
$ protoc -I=protos --go_out=internal --go-grpc_out=internal --grpc-gateway_out=internal/protos --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative  protos/*.proto
```

**Generate node protocol buffers files**
```sh
# install grpc-tools
$ npm install -g grpc-tools
```

```sh
# generate node files
$ grpc_tools_node_protoc -I ./protos --js_out=import_style=commonjs,binary:./sdk/js --grpc_out=grpc_js:./sdk/js protos/*.proto
```

## Contributing
1. Clone it (`git clone git@github.com:jasonsites/gosk-grpc.git`)
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Commit your changes using [conventional changelog standards](https://www.conventionalcommits.org) (`git commit -m 'feat(ENG-1234): adds my new feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Ensure linting and tests are all passing
1. Create new Pull Request

## Releasing
*Requires [standard-version](https://github.com/conventional-changelog/standard-version)*
1. Merge fixes & features to master
1. Run `npx standard-version`
1. Push release & release tag to github (`git push --follow-tags origin master`)

## License
Copyright (c) 2022 Jason Sites
