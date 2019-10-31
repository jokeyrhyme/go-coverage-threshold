package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoPath(t *testing.T) {
	t.Parallel()

	got, module, err := goPath()
	assert.NotEmpty(t, got)
	assert.Nil(t, err)
	// until this project moves to modules - this should be empty
	assert.Empty(t, module)
}
