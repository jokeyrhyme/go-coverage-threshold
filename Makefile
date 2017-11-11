ci: lint test test-cover test-race

lint: setup-lint
	gometalinter --concurrency 2 --deadline 5m --tests --vendor ./...

run:
	go run ./cmd/go-coverage-threshold/main.go -threshold 0

setup-lint:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

test:
	go test ./...

test-cover:
	go run ./cmd/go-coverage-threshold/main.go -threshold 0

test-race:
	go test -race ./...

.DO_NOT_CACHE: lint test test-cover test-race
