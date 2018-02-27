# go-coverage-threshold [![Build Status](https://travis-ci.org/jokeyrhyme/go-coverage-threshold.svg?branch=master)](https://travis-ci.org/jokeyrhyme/go-coverage-threshold)

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

## Configuration

You may place a .cover.toml file at the root of your project,
as an alternative to using command line arguments, e.g:

```toml
# important: specify "50.0" if you want 50% coverage,
# "50" without the ".0" will not work
threshold = 50.0
```

Note that command line arguments take precedence over configuration files

Note that .cover.toml files in sub-directories take precedence over parent directories,
so you may have a threshold for the whole project as a rule,
yet define exceptions for certain sub-directories, e.g:

* PROJECT_ROOT/.cover.toml: threshold = 80.0
* PROJECT_ROOT/cmd/.cover.toml: threshold = 10.0
