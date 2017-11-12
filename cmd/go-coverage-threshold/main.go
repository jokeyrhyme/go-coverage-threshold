package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jokeyrhyme/go-coverage-threshold/pkg/cover"
)

const (
	thresholdDefault = 80.0
	thresholdUsage   = "threshold that coverage must exceed"
)

var (
	threshold float64
)

func flags() {
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

	var config *cover.Config
	if len(os.Args) < 2 {
		// user did not specify -t or -threshold
		config, err = cover.Load("")
		if err != nil {
			fmt.Printf("no arguments specified, and unable to load .cover.toml file: %v\n", err)
			config = &cover.Config{
				Threshold: thresholdDefault,
			}
		}
	} else {
		config = &cover.Config{
			Threshold: threshold,
		}
	}

	exitCode := 0
	for _, e := range cover.Parse(output) {
		e.Threshold = config.Threshold
		if e.Failed() {
			exitCode = 1
		}
		fmt.Println(e.String())
	}
	os.Exit(exitCode)
}
