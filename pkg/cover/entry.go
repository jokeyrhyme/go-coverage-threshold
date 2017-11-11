package cover

import (
	"strconv"
	"strings"
)

// Entry represents a per-path coverage measurement.
type Entry struct {
	Coverage  float64
	Duration  string // TODO: parse this further into a time.Duration
	Path      string
	Status    string
	Threshold float64
}

// Failed checks if the threshold was not met.
func (e *Entry) Failed() bool {
	return e.Coverage < e.Threshold
}

// Passed checks if the threshold was met.
func (e *Entry) Passed() bool {
	return e.Coverage >= e.Threshold
}

func (e *Entry) String() string {
	var status string
	if e.Passed() {
		status = " "
	} else if e.Failed() {
		status = "x"
	}
	return strings.Join([]string{
		status,
		e.Path,
		e.Duration,
		"coverage: " + strconv.FormatFloat(e.Coverage, 'f', 2, 64) + "% of statements",
		"(threshold: " + strconv.FormatFloat(e.Threshold, 'f', 2, 64) + "%)",
	}, " ")
}
