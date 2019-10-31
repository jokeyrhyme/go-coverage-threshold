package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/jokeyrhyme/go-coverage-threshold/pkg/cover"
)

const (
	thresholdDefault = 80.0
	thresholdUsage   = "threshold that coverage must exceed"
)

var (
	threshold float64
)

func config(s string) *cover.Config {
	if len(os.Args) >= 2 {
		// user specified -t or -threshold
		// args take precedence over .cover.toml files
		return &cover.Config{
			Threshold: threshold,
		}
	}

	config, err := cover.Load(s)
	if err != nil {
		fmt.Printf("no arguments specified, and unable to load .cover.toml file: %v\n", err)
		return &cover.Config{
			Threshold: thresholdDefault,
		}
	}
	return config
}

func flags() {
	flag.Float64Var(&threshold, "threshold", thresholdDefault, thresholdUsage)
	flag.Float64Var(&threshold, "t", thresholdDefault, thresholdUsage+" (shorthand)")
	flag.Parse()
}

func goPath() (string, string, error) {
	gopath, ok := os.LookupEnv("GOPATH")
	if ok {
		return gopath, "", nil
	}

	// check if the project uses go modules by running go mod why
	// for a project that does not use go modules this command will fail
	cmd := exec.Command("go", "mod", "why") // nolint: gas,gosec
	if out, err := cmd.CombinedOutput(); err == nil {
		// go modules are active - use current working directory
		moduleOut := strings.Split(string(out), "\n")
		moduleName := moduleOut[1]

		// look for working directory - this is non optional
		// panic if no PWD is not set
		if pwd, ok := os.LookupEnv("PWD"); ok {
			return pwd, moduleName, nil
		}
		log.Fatalf("PWD is not set in ENV, when using modules it is necessary to know current working directory in order to build package path")
	}

	home, ok := os.LookupEnv("HOME")
	if !ok {
		return "", "", errors.New("no GOPATH or HOME in environment")
	}
	stat, err := os.Stat(path.Join(home, "go", "src"))
	if err != nil || !stat.IsDir() {
		return "", "", errors.New("$HOME/go is not a valid GOPATH")
	}
	return path.Join(home, "go"), "", nil
}

func main() {
	flags()

	output, err := cover.Run()
	if err != nil {
		os.Exit(1)
	}

	gp, module, err := goPath()
	if err != nil {
		os.Exit(1)
	}

	exitCode := 0
	for _, e := range cover.Parse(output) {
		realPath := ""
		if len(module) == 0 {
			realPath = path.Join(gp, "src", e.Path)
		} else {
			realPath = path.Join(gp, strings.ReplaceAll(e.Path, module, ""))
		}
		cfg := config(realPath)

		e.Threshold = cfg.Threshold

		if e.Failed() {
			exitCode = 1
		}
		fmt.Println(e.String())
	}
	os.Exit(exitCode)
}
