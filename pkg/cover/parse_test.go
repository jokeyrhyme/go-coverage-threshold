package cover_test

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/jokeyrhyme/go-coverage-threshold/pkg/cover"
	"github.com/stretchr/testify/assert"
)

func readFixture(t *testing.T, s string) []byte {
	f, err := os.Open(path.Join("fixtures", s))
	if err != nil {
		t.Fatalf("unable to open fixture: %v\n", err)
	}
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			t.Fatalf("unable to close fixture: %v\n", closeErr)
		}
	}()

	buffer := &bytes.Buffer{}
	_, err = buffer.ReadFrom(f)
	if err != nil {
		t.Fatalf("unable to read fixture into buffer: %v\n", err)
	}

	return buffer.Bytes()
}

func TestParse(t *testing.T) {
	t.Parallel()

	type Case struct {
		fixtureFile string
		want        []*cover.Entry
	}
	cases := []Case{
		{
			fixtureFile: "go-1-9-2-test-cover-0-nil.txt",
			want: []*cover.Entry{
				{
					Coverage:  0.0,
					Duration:  "0.001s",
					Path:      "github.com/jokeyrhyme/go-coverage-threshold/cmd/go-coverage-threshold",
					Status:    "ok",
					Threshold: 0.0,
				},
				{
					Coverage:  0.0,
					Duration:  "",
					Path:      "github.com/jokeyrhyme/go-coverage-threshold/pkg/cover",
					Status:    "?",
					Threshold: 0.0,
				},
			},
		},
	}

	for i, c := range cases {
		func(i int, c Case) {
			t.Run(fmt.Sprintf("cases[%d]", i), func(t *testing.T) {
				t.Parallel()

				got := cover.Parse(readFixture(t, c.fixtureFile))
				assert.Equal(t, c.want, got)
			})
		}(i, c)
	}
}
