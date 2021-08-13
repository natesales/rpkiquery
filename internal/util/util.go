package util

import (
	"encoding/hex"
	"strings"
)

// FormatHex formats a byte slice into a colon separated hex string
func FormatHex(b []byte) string {
	buf := make([]byte, 0, 3*len(b))
	x := buf[1*len(b) : 3*len(b)]
	hex.Encode(x, b)
	for i := 0; i < len(x); i += 2 {
		buf = append(buf, x[i], x[i+1], ':')
	}
	if len(buf) > 1 {
		return strings.ToUpper(string(buf[:len(buf)-1]))
	}
	return string(b)
}
