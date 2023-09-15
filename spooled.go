package spooledtempfile

import (
	"io"
	"io/fs"
)

type SpooledTemporaryFile struct {
	// sizeMax is the max allowed size for the in-memory buffer
	// if the content becomes larger than that value, it will be
	// dumped into a temporary file.
	sizeMax int
	// sizeWrote keeps track of the number of bytes that have been
	// written to the buffer OR the file.
	sizeWrote int
	// rolledOver keeps track if the content was dumped to file.
	rolledOver bool
	buffer     []byte
	file       fs.File

	io.Reader
	io.Writer
	io.Closer
}

func NewSpooledTemporaryFile(maxSize int, buffer []byte) *SpooledTemporaryFile {
	return &SpooledTemporaryFile{
		sizeMax: maxSize,
		buffer:  buffer,
	}
}

// Write appends the provided buffer into the internal buffer OR
// may directly create a file if the size of the provided buffer
// is larger than the max allowed size.
func (s *SpooledTemporaryFile) Write(bytes []byte) (int, error) {
	return 0, nil
}

// Read copies the content of the internal buffer OR the file into
// the provided buffer.
func (s *SpooledTemporaryFile) Read(bytes []byte) (int, error) {
	return 0, nil
}

// Close will close the underlying file object. It is NOOP
// when the buffer is still in-memory.
func (s *SpooledTemporaryFile) Close() error {
	return nil
}

// Rollover will write the buffer to file even if its
// smaller than sizeMax.
func (s *SpooledTemporaryFile) Rollover() error {
	s.rolledOver = true
	return nil
}
