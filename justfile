set shell := ["bash", "-uc"]

# Defaults ========================================================================================
project := 'domain'

# Commands ========================================================================================
# show this help
help:
	just --list

# remove build related files
clean:
  rm -rf bin
  rm -rf out
  rm -f profile.cov

# Lint ============================================================================================
# lint all
# lint: TODO

# Migrations ======================================================================================
# migrate down
migrate-down db +step='-all':
  migrate -path ./internal/migrations -database postgres://postgres:postgres@host:5432/{{db}}?sslmode=disable down {{step}}

# migrate up
migrate-up db *step:
  migrate -path ./internal/migrations -database postgres://postgres:postgres@host:5432/{{db}}?sslmode=disable up {{step}}

# migrate up -all (alias)
migrate:
  just migrate-up svcdb

# create migration with {{name}}
migrate-create name:
	migrate create -ext sql -dir ./internal/migrations -format unix {{name}}

# Run =============================================================================================
# run the test client
client:
	go run cmd/client/client.go

# run the server in dev mode
server:
  air

# Test ============================================================================================
# run tests
test:
  just migrate-up testdb
  gotestsum -- -v -race ./...

# run tests with coverage report
coverage:
  just migrate-up testdb
  gotestsum --jsonfile test-output.log --junitfile junit.xml -- -coverpkg=$(go list ./... | grep -v proto | grep -v testutils | tr '\n' ',') -covermode=count -coverprofile=profile.cov ./...
  go tool cover -o coverage.html -html=profile.cov