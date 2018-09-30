package sfp

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		raw, enc []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte{'a'}, []byte{'a'}},
		{[]byte{delim}, []byte{escape, delim ^ xor}},
		{[]byte{escape}, []byte{escape, escape ^ xor}},
		{[]byte{delim, escape}, []byte{escape, delim ^ xor, escape, escape ^ xor}},
		{[]byte{delim, delim}, []byte{escape, delim ^ xor, escape, delim ^ xor}},
		{[]byte{escape, escape}, []byte{escape, escape ^ xor, escape, escape ^ xor}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Encode % x", tc.raw), func(t *testing.T) {
			res := Encode(tc.raw)
			if !bytes.Equal(res, tc.enc) {
				t.Errorf("got %q, want %q", res, tc.enc)
			}
		})
		t.Run(fmt.Sprintf("Decode % x", tc.enc), func(t *testing.T) {
			res := Decode(tc.enc)
			if !bytes.Equal(res, tc.raw) {
				t.Errorf("got %q, want %q", res, tc.raw)
			}
		})
	}
}
