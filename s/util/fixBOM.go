package util

import (
	"bytes"
)

var utf8bom = []byte{239, 187, 191}

// utf8のbomを消す
func NormalizeUTF8BOM(in []byte) []byte {
	return bytes.TrimPrefix(in, utf8bom)
}
