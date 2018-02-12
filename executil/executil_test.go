package executil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunLines(t *testing.T) {
	assert.NoError(t, RunLines(func(line string) {
		assert.Equal(t, "test", line)
	}, "echo", "-n", "test"), "Last line without newline")

	assert.NoError(t, RunLines(func(line string) {
		assert.Equal(t, "test", line)
	}, "echo", "test"), "Last line with newline")

	got := make([]string, 0)
	assert.NoError(t, RunLines(func(line string) {
		got = append(got, line)
	}, "echo", "1\n2"), "Last line with newline")
	assert.Equal(t, []string{"1", "2"}, got)
}
