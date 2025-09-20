# spooled-temporary-file-go
this library may be used by projects that parse large payloads as well as small ones, so if the payload size is larger than the specificed size, it will be written to a temporary file, otherwise kept in memory. This essentially reduces the memory footprint of a cloud proxy dealing with large payloads.

Inspired by https://docs.python.org/3/library/tempfile.html#tempfile.SpooledTemporaryFile but for Go
