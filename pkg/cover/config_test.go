package cover

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("unable to determine working directory: %v\n", err)
	}

	type Case struct {
		cwd     string
		want    *Config
		wantErr error
	}
	cases := []Case{
		{
			cwd:     wd,
			want:    &Config{70.0},
			wantErr: nil,
		},
		{
			cwd:     "/",
			want:    nil,
			wantErr: errors.New("file not found"),
		},
	}

	for i, c := range cases {
		func(i int, c Case) {
			t.Run(fmt.Sprintf("cases[%d]", i), func(t *testing.T) {
				t.Parallel()

				got, err := Load(c.cwd)
				assert.Equal(t, c.want, got)
				assert.Equal(t, c.wantErr, err)
			})
		}(i, c)
	}
}
