ci: lint test test-cover test-race

lint: setup-lint vendor
	gometalinter --concurrency 2 --deadline 5m --exclude libexec --tests --vendor ./...

run: vendor
	go run ./cmd/go-coverage-threshold/main.go -threshold 0

setup-vendor:
	go get -u github.com/golang/dep/cmd/dep

setup-lint:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

test: vendor
	go test ./...

test-cover: vendor
	go run ./cmd/go-coverage-threshold/main.go -threshold 0

test-race: vendor
	go test -race ./...

vendor: setup-vendor
	dep ensure

.DO_NOT_CACHE: lint test test-cover test-race
