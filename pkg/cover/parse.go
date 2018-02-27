package cover

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
)

var (
	okRegexp  = regexp.MustCompile(`^(?P<status>ok)\s+(?P<path>\S+)\s+(?P<duration>\d+\.\d+\w|\(cached\))\s+coverage: (?P<coverage>\d+\.\d+)% of statements$`)
	nilRegexp = regexp.MustCompile(`^(?P<status>\?)\s+(?P<path>\S+).+$`)
)

// Parse the stdout result from `go test -cover ./...`.
func Parse(result []byte) []*Entry {
	scanner := bufio.NewScanner(bytes.NewReader(result))
	scanner.Split(onEOL)

	var entries []*Entry
	for scanner.Scan() {
		line := scanner.Text()

		matches := mergeMaps(namedCaptures(okRegexp, line), namedCaptures(nilRegexp, line))
		if len(matches) == 0 {
			continue
		}

		entry := &Entry{
			Path:   matches["path"],
			Status: matches["status"],
		}
		if duration, ok := matches["duration"]; ok {
			entry.Duration = duration
		}
		if coverage, ok := matches["coverage"]; ok {
			parsed, err := strconv.ParseFloat(coverage, 64)
			if err == nil {
				entry.Coverage = parsed
			}
		}

		entries = append(entries, entry)
	}
	return entries
}

func onEOL(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			return i + 1, data[:i], nil
		}
	}
	return 0, data, bufio.ErrFinalToken
}
