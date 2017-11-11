package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jokeyrhyme/go-coverage-threshold/pkg/cover"
)

var (
	threshold float64
)

func flags() {
	const (
		thresholdDefault = 80.0
		thresholdUsage   = "threshold that coverage must exceed"
	)
	flag.Float64Var(&threshold, "threshold", thresholdDefault, thresholdUsage)
	flag.Float64Var(&threshold, "t", thresholdDefault, thresholdUsage+" (shorthand)")
	flag.Parse()
}

func main() {
	flags()

	output, err := cover.Run()
	if err != nil {
		os.Exit(1)
	}

	exitCode := 0
	for _, e := range cover.Parse(output) {
		e.Threshold = threshold
		if e.Failed() {
			exitCode = 1
		}
		fmt.Println(e.String())
	}
	os.Exit(exitCode)
}
