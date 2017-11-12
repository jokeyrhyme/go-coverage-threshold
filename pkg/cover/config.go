package cover

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

const (
	configFile = ".cover.toml"
)

// Config holds settings for coverage and thresholds.
type Config struct {
	Threshold float64 `toml:"threshold"`
}

// Load configuration values from files.
func Load(wd string) (*Config, error) {
	cwd := wd
	if len(cwd) == 0 {
		var err error
		cwd, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}

	which, err := findUp(cwd, configFile)
	if err != nil {
		return nil, err
	}

	bytes, err := readFile(which)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = toml.Unmarshal(bytes, config)
	if err != nil {
		fmt.Printf("bad TOML? %s: %v\n", which, err)
		return nil, err
	}

	return config, nil
}

func readFile(s string) ([]byte, error) {
	f, err := os.Open(s)
	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			fmt.Printf("unable to close config: %v\n", closeErr)
		}
	}()

	buffer := &bytes.Buffer{}
	_, err = buffer.ReadFrom(f)

	return buffer.Bytes(), err
}
