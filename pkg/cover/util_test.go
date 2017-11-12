package cover

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUp(t *testing.T) {
	t.Parallel()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("unable to determine working directory: %v\n", err)
	}

	type Case struct {
		cwd     string
		s       string
		want    string
		wantErr error
	}
	cases := []Case{
		{
			cwd:     wd,
			s:       "util_test.go",
			want:    path.Join(wd, "util_test.go"),
			wantErr: nil,
		},
		{
			cwd:     path.Join(wd, "fixtures"),
			s:       "util_test.go",
			want:    path.Join(wd, "util_test.go"),
			wantErr: nil,
		},
		{
			cwd:     wd,
			s:       "really_really_should_not_exist_" + string(rand.Int()),
			want:    "",
			wantErr: errors.New("file not found"),
		},
	}

	for i, c := range cases {
		func(i int, c Case) {
			t.Run(fmt.Sprintf("cases[%d]", i), func(t *testing.T) {
				t.Parallel()

				got, err := findUp(c.cwd, c.s)
				assert.Equal(t, c.want, got)
				assert.Equal(t, c.wantErr, err)
			})
		}(i, c)
	}
}

func TestMergeMaps(t *testing.T) {
	t.Parallel()

	type Case struct {
		maps []map[string]string
		want map[string]string
	}
	cases := []Case{
		{
			maps: []map[string]string{},
			want: map[string]string{},
		},
		{
			maps: []map[string]string{
				map[string]string{
					"abc": "123",
					"def": "456",
				},
				map[string]string{
					"def": "789",
					"ghi": "101112",
					"jkl": "131415",
				},
				map[string]string{
					"jkl": "161718",
				},
			},
			want: map[string]string{
				"abc": "123",
				"def": "789",
				"ghi": "101112",
				"jkl": "161718",
			},
		},
	}

	for i, c := range cases {
		func(i int, c Case) {
			t.Run(fmt.Sprintf("cases[%d]", i), func(t *testing.T) {
				t.Parallel()

				got := mergeMaps(c.maps...)
				assert.Equal(t, c.want, got)
			})
		}(i, c)
	}
}

func TestNamedCaptures(t *testing.T) {
	t.Parallel()

	type Case struct {
		re   *regexp.Regexp
		s    string
		want map[string]string
	}
	cases := []Case{
		{
			re:   regexp.MustCompile(""),
			s:    "",
			want: map[string]string{},
		},
		{
			re: regexp.MustCompile(`(?P<nums>\d+)`),
			s:  "abc123",
			want: map[string]string{
				"nums": "123",
			},
		},
	}

	for i, c := range cases {
		func(i int, c Case) {
			t.Run(fmt.Sprintf("cases[%d]", i), func(t *testing.T) {
				t.Parallel()

				got := namedCaptures(c.re, c.s)
				assert.Equal(t, c.want, got)
			})
		}(i, c)
	}
}
