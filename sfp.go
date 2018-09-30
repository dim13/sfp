// Package sfp implements RFC1662 Serial Framing Protocol
package sfp

import "bytes"

const (
	delim  = 0x7d
	escape = 0x7e
	xor    = 0x20
)

// Encode frame
func Encode(frame []byte) []byte {
	buf := new(bytes.Buffer)
	for _, v := range frame {
		switch v {
		case delim, escape:
			buf.WriteByte(escape)
			buf.WriteByte(v ^ xor)
		default:
			buf.WriteByte(v)
		}
	}
	return buf.Bytes()
}

// Decode frame
func Decode(frame []byte) []byte {
	buf := new(bytes.Buffer)
	inEscape := false
	for _, v := range frame {
		switch {
		case !inEscape && v == escape:
			inEscape = true
		case inEscape:
			buf.WriteByte(v ^ xor)
			inEscape = false
		default:
			buf.WriteByte(v)
		}
	}
	return buf.Bytes()
}
