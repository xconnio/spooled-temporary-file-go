# spooled-file-go
this library may be used by projects that parse large payloads as well as small ones, so if the payload size is larger than a specific file, it will be written to a temporary file, otherwise kept in memory. This essentially reduces the memory footprint of an API server OR a websocket server

Inspired by https://docs.python.org/3/library/tempfile.html#tempfile.SpooledTemporaryFile but for Go
