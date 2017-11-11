# go-coverage-threshold

keep your per-path Go test coverage above a threshold


## Installation

```sh
go get -u github.com/jokeyrhyme/go-coverage-threshold/cmd/go-coverage-threshold
```


## Usage

```
$ go-coverage-threshold --help

Usage of go-coverage-threshold:
  -t float
        threshold that coverage must exceed (shorthand) (default 80)
  -threshold float
        threshold that coverage must exceed (default 80)
```

`go-coverage-threshold` will internally execute `go test -cover ./...` for you

If any of your `./...` paths (paths that contain .go files) have a test coverage percentage that is below the threshold,
then it exits with a non-zero exit code

This is useful for Continuous Integration workflows where you want to maintain and encourage test coverage


## Roadmap

-   [ ] read threshold from .cover.toml file in project root

-   [ ] optional per-path thresholds
