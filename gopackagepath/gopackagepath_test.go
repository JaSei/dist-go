package gopackagepath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	gpp, err := New("github.com")
	assert.Error(t, err)

	gpp, err = New("github.com/JaSei")
	assert.Error(t, err)

	gpp, err = New("github.com/JaSei/test")
	assert.NoError(t, err)

	assert.Equal(t, "github.com", gpp.Repo())
	assert.Equal(t, "JaSei", gpp.User())
	assert.Equal(t, "test", gpp.Package())
	assert.Equal(t, "JaSei/test", gpp.UserPackage())
	assert.Equal(t, "github.com/JaSei/test", gpp.FullPackage())

	gpp, err = New("github.com/JaSei/test/testify/assert")
	assert.NoError(t, err)
	assert.Equal(t, "testify/assert", gpp.SubPackage())
}
