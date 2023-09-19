package spooledtempfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpooledTemporaryFile_Write(t *testing.T) {
	stf := NewSpooledTemporaryFile(30, nil)

	testData := []byte("Hello World!")
	n, err := stf.Write(testData)
	assert.NoError(t, err)
	assert.Equal(t, len(testData), n)

	// Check the sizeWrote field.
	assert.Equal(t, n, stf.sizeWrote)

	// Attempt to write more than the max size.
	testData = []byte("This is too much data!")
	n, err = stf.Write(testData)
	assert.Equal(t, len(testData), n)

	// Check that it's rolled over after reaching the max size.
	assert.True(t, stf.rolledOver)
}
