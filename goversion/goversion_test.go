package goversion

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
)

func TestGoVersion(t *testing.T) {
	vers, err := GoVersion()
	assert.NoError(t, err)
	sRange, err := semver.ParseRange(">=1.5.0 <2.0.0")
	assert.True(t, sRange(*vers))

	func() {
		command = []string{"error"}
		defer func() { command = []string{"version"} }()
		vers, err = GoVersion()
		assert.Error(t, err)
		assert.Nil(t, vers)
	}()
}
