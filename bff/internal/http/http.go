package http

import (
	"bytes"
	"io"
)

func ReadCloserToString(rc io.ReadCloser) string {
	if rc == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(rc)
	return buf.String()
}
