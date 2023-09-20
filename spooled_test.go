package spooledtempfile_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xconnio/spooled-temporary-file"
)

func TestSpooledTemporaryFile_Write(t *testing.T) {
	stf := spooledtempfile.NewSpooledTemporaryFile(30, nil)

	testData := []byte("Hello World!")
	n, err := stf.Write(testData)
	assert.NoError(t, err)
	assert.Equal(t, len(testData), n)

	// Check the sizeWrote field.
	assert.Equal(t, n, stf.SizeWrote())

	// Attempt to write more than the max size.
	testData = []byte("This is too much data!")
	n, _ = stf.Write(testData)
	assert.Equal(t, len(testData), n)

	// Check that it's rolled over after reaching the max size.
	assert.True(t, stf.RolledOver())
}

func TestSpooledTemporaryFile_Read(t *testing.T) {
	stf := spooledtempfile.NewSpooledTemporaryFile(1024, nil)

	data := []byte("Hello World!")

	_, err := stf.Write(data)
	require.NoError(t, err)

	readBuffer := make([]byte, len(data))

	n, err := stf.Read(readBuffer)
	require.NoError(t, err)
	require.Equal(t, len(data), n)
	require.Equal(t, data, readBuffer)
}
